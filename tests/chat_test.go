package tests

import (
	"bytes"
	"encoding/json"
	"livechat-support/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestSaveMessage(t *testing.T) {
	utils.InitTestLogger() // Initialize logger for test
	router := setupRouter()

	message := map[string]string{
		"username": "testuser",
		"message":  "Hello, this is a test message!",
	}
	jsonMessage, _ := json.Marshal(message)

	req, _ := http.NewRequest("POST", "/save-message", bytes.NewBuffer(jsonMessage))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	utils.Logger.Info("Testing SaveMessage Endpoint",
		zap.Int("StatusCode", w.Code),
		zap.String("Response", w.Body.String()),
	)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRecentMessages(t *testing.T) {
	utils.InitTestLogger()
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/recent-messages", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	utils.Logger.Info("Testing GetRecentMessages Endpoint",
		zap.Int("StatusCode", w.Code),
		zap.String("Response", w.Body.String()),
	)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWebSocketConnection(t *testing.T) {
	utils.InitTestLogger()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				utils.Logger.Error("WebSocket receive error", zap.Error(err))
				return
			}

			utils.Logger.Info("WebSocket received message", zap.String("Message", msg))

			// Echo message back
			err = websocket.Message.Send(ws, msg)
			if err != nil {
				utils.Logger.Error("WebSocket send error", zap.Error(err))
			}
		}).ServeHTTP(w, r)
	}))
	defer server.Close()

	// Connect to the WebSocket server
	ws, err := websocket.Dial("ws://"+server.Listener.Addr().String()+"/ws", "", "http://localhost/")
	assert.NoError(t, err)
	defer ws.Close()

	// Send test message
	testMessage := "Hello WebSocket!"
	err = websocket.Message.Send(ws, testMessage)
	assert.NoError(t, err)

	// Receive echoed message
	var receivedMessage string
	err = websocket.Message.Receive(ws, &receivedMessage)
	assert.NoError(t, err)

	utils.Logger.Info("Testing WebSocket Connection",
		zap.String("SentMessage", testMessage),
		zap.String("ReceivedMessage", receivedMessage),
	)

	assert.Equal(t, testMessage, receivedMessage)
}
