package main

import (
	"shopsocial-backend/api"
	"shopsocial-backend/config"
	"shopsocial-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize logger
	logger.InitLogger()
	defer logger.Log.Sync() // Flush logs on exit

	// Connect to MongoDB
	config.ConnectDB()

	// Setup Gin
	router := gin.Default()

	// Register all routes
	api.RegisterRoutes(router)

	// Start the server
	port := ":8080"
	router.Run(port)
	logger.Log.Info("Starting server", zap.String("port", port))
}
