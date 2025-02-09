package main

import (
	"fmt"
	"livechat-support/config"
	"livechat-support/database"
	"livechat-support/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func main() {
	// Load Configuration
	fmt.Println("Loading Configuration...")
	config.LoadConfig()

	// Initialize Database
	fmt.Println("Initializing Database...")
	database.ConnectDB()

	// Initialize Redis
	fmt.Println("Initializing Redis...")
	utils.ConnectRedis()

	// Initialize Logger
	fmt.Println("Initializing Logger...")
	utils.InitLogger()

	// Create Gin router
	r := gin.Default()

	// Register Routes
	routes.RegisterRoute(r)

	// Start Server
	fmt.Println("Starting Server...")
	utils.Logger.Info("Server is running on port 8080")
	r.Run(":8080")
}
