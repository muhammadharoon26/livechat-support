package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"livechat-support/config"
	"livechat-support/controllers"
	"livechat-support/database"
	"livechat-support/middleware"
	"livechat-support/models"
	"livechat-support/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	router *gin.Engine
	db     *gorm.DB
}

func (s *TestSuite) SetupSuite() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize test database
	var err error
	s.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		s.T().Fatal("Failed to connect to test database:", err)
	}

	// Migrate the schema
	s.db.AutoMigrate(&models.User{}, &models.Message{})

	// Set the global DB instance
	database.DB = s.db

	// Initialize Redis for testing
	utils.ConnectRedis()

	// Initialize test logger
	utils.InitTestLogger()

	// Setup router with all required routes
	s.router = gin.New()
	s.router.POST("/register", controllers.Register)
	s.router.POST("/login", controllers.Login)
	s.router.POST("/save-message", controllers.SaveMessage)
	s.router.GET("/recent-messages", controllers.GetRecentMessages)
}

func (s *TestSuite) TearDownSuite() {
	// Clean up database
	db, err := s.db.DB()
	if err == nil {
		db.Close()
	}

	// Clean up Redis
	utils.RDB.FlushAll(context.Background())
	utils.RDB.Close()
}

func (s *TestSuite) SetupTest() {
	// Clear database before each test
	s.db.Exec("DELETE FROM users")
	s.db.Exec("DELETE FROM messages")

	// Clear Redis before each test
	utils.RDB.FlushAll(context.Background())
}

// Test user registration
func (s *TestSuite) TestRegister() {
	body := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusCreated, w.Code)

	// Verify user was created
	var user models.User
	result := s.db.Where("username = ?", "testuser").First(&user)
	assert.Nil(s.T(), result.Error)
	assert.NotEmpty(s.T(), user.ID)

	// Verify password was hashed
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("testpassword"))
	assert.Nil(s.T(), err)
}

// Test user login
func (s *TestSuite) TestLogin() {
	// Create a test user first
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	user := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	}
	s.db.Create(&user)

	body := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
}

// Test message saving
func (s *TestSuite) TestSaveMessage() {
	// Create a test user first
	user := models.User{
		Username: "testuser",
		Password: "hashedpass",
	}
	s.db.Create(&user)

	// Create the message
	message := models.Message{
		UserID:  user.ID,
		Content: "Test message",
	}
	jsonMessage, _ := json.Marshal(message)

	// Create the request
	req := httptest.NewRequest("POST", "/save-message", bytes.NewBuffer(jsonMessage))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Send the request
	s.router.ServeHTTP(w, req)

	// Check response status
	assert.Equal(s.T(), http.StatusOK, w.Code)

	// Verify message was saved in database
	var savedMessage models.Message
	result := s.db.Where("user_id = ?", user.ID).First(&savedMessage)
	assert.NoError(s.T(), result.Error)
	assert.Equal(s.T(), "Test message", savedMessage.Content)

	// Verify message was saved in Redis
	ctx := context.Background()
	messages, err := utils.RDB.LRange(ctx, "recent_messages", 0, -1).Result()
	assert.NoError(s.T(), err)
	assert.Contains(s.T(), messages, "Test message")
}

// Test getting recent messages
func (s *TestSuite) TestGetRecentMessages() {
	ctx := context.Background()

	// Add some test messages to Redis
	testMessages := []string{"Message 1", "Message 2", "Message 3"}
	for _, msg := range testMessages {
		utils.RDB.LPush(ctx, "recent_messages", msg)
	}

	req := httptest.NewRequest("GET", "/recent-messages", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)

	var response map[string][]string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(s.T(), 3, len(response["messages"]))
	assert.Equal(s.T(), "Message 3", response["messages"][0])
}

// Test JWT token generation
func (s *TestSuite) TestJWTToken() {
	user := models.User{
		Username: "testuser",
		Password: "hashedpass",
	}
	s.db.Create(&user)

	token := middleware.GenerateToken(user)
	assert.NotEmpty(s.T(), token)

	// Verify token expiration
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	assert.Nil(s.T(), err)

	exp := int64(claims["exp"].(float64))
	assert.Greater(s.T(), exp, time.Now().Unix())
	assert.Less(s.T(), exp, time.Now().Add(25*time.Hour).Unix())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
