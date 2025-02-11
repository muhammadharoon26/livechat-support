package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var msg string
		if err := conn.ReadJSON(&msg); err != nil {
			delete(clients, conn)
			break
		}
		for client := range clients {
			client.WriteJSON(msg)
		}
	}
}
