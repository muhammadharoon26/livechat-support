package tests

import (
	"bytes"
	"encoding/json"
	"livechat-support/database"
	"livechat-support/models"
	"livechat-support/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

func TestRegister(t *testing.T) {
	database.ConnectDB()
	r := setupTestRouter()

	user := models.User{
		Username: "testuser",
		Password: "password123",
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLogin(t *testing.T) {
	database.ConnectDB()
	r := setupTestRouter()

	user := models.User{
		Username: "testuser",
		Password: "password123",
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code) // Should be unauthorized unless user exists
}
