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

func setupChatRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

func TestSendMessage(t *testing.T) {
	database.ConnectDB() // Ensure DB is connected
	router := setupChatRouter()

	messageData := map[string]string{
		"sender":  "testuser",
		"message": "Hello, this is a test message!",
	}

	jsonData, _ := json.Marshal(messageData)
	req, _ := http.NewRequest("POST", "/send-message", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Message sent successfully")
}

func TestReceiveMessages(t *testing.T) {
	router := setupChatRouter()

	req, _ := http.NewRequest("GET", "/messages", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "messages")
}
