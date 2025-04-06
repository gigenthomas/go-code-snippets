package logger

import (
	"log"

	"go.uber.org/zap"
)

// InitLogger initializes and replaces the global logger
func InitLogger() *zap.Logger {
	// Initialize the production-level logger (you can choose development mode here too)
	logger, err := zap.NewProduction() // or zap.NewDevelopment() for dev logs
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	// Replace the global logger with the custom logger
	zap.ReplaceGlobals(logger)
	return logger
}
