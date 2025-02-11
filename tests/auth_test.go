package tests

import (
	"bytes"
	"encoding/json"
	"livechat-support/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRegister(t *testing.T) {
	utils.InitTestLogger() // Initialize logger for tests
	router := setupRouter()

	body := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	utils.Logger.Info("Testing Register Endpoint",
		zap.Int("StatusCode", w.Code),
		zap.String("Response", w.Body.String()),
	)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLogin(t *testing.T) {
	utils.InitTestLogger()
	router := setupRouter()

	body := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	utils.Logger.Info("Testing Login Endpoint",
		zap.Int("StatusCode", w.Code),
		zap.String("Response", w.Body.String()),
	)

	assert.Equal(t, http.StatusOK, w.Code)
}
