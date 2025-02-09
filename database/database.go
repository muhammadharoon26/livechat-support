package database

import (
	"livechat-support/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	// AutoMigrate models
	DB.AutoMigrate(&models.User{}, &models.Message{})
}
