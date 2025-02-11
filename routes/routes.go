package routes

import (
	"livechat-support/controllers"
	"livechat-support/websocket"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/ws", websocket.HandleConnections)
	r.POST("/save-message", controllers.SaveMessage)
	r.GET("/recent-messages", controllers.GetRecentMessages)
}
