package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"livechat-support/database"
	"livechat-support/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

func TestRegisterUser(t *testing.T) {
	database.ConnectDB() // Ensure DB is connected
	router := setupRouter()

	userData := map[string]string{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "securepassword",
	}

	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()

	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "securepassword",
	}

	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestProtectedRouteWithoutToken(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}
