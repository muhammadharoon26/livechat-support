package controllers

import (
	"livechat-support/database"
	"livechat-support/models"
	"livechat-support/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&msg)
	utils.RDB.LPush(c, "recent_messages", msg.Content)
	c.JSON(http.StatusOK, gin.H{"message": "Message saved successfully"})
}

func GetRecentMessages(c *gin.Context) {
	messages, _ := utils.RDB.LRange(c, "recent_messages", 0, 10).Result()
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
