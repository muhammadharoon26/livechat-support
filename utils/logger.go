package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	logger, _ := zap.NewProduction()
	Logger = logger
	defer logger.Sync() // Flush logs before exiting
}

func InitTestLogger() {
	logger, _ := zap.NewDevelopment() // Use a more readable format for tests
	Logger = logger
	defer logger.Sync()
}
