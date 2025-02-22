package logger

import (
	"go.uber.org/zap"
	"os"
)

var Log *zap.Logger

// InitLogger initializes the zap logger
func InitLogger() {
	var err error

	// Check environment mode
	if os.Getenv("ENV") == "production" {
		// Production: JSON logging
		Log, err = zap.NewProduction()
	} else {
		// Development: Human-readable logs
		Log, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	Log.Info("Logger initialized successfully!")
}
