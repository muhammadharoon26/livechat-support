package tests

import (
	"bytes"
	"encoding/json"
	"livechat-support/database"
	"livechat-support/models"
	"livechat-support/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveMessage(t *testing.T) {
	database.ConnectDB()
	utils.ConnectRedis()
	r := setupTestRouter()

	message := models.Message{
		UserID:  1,
		Content: "Hello, world!",
	}
	body, _ := json.Marshal(message)

	req, _ := http.NewRequest("POST", "/save-message", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRecentMessages(t *testing.T) {
	utils.ConnectRedis()
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/recent-messages", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
