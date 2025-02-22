package main

import (
	"net/http"
	"shopsocial-backend/api"
	"shopsocial-backend/config"
	"shopsocial-backend/pkg/logger"
	"shopsocial-backend/pkg/responses"

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

	router.NoRoute(func(c *gin.Context) {
		// Return a standard JSON error for all invalid routes
		responses.SendError(c, http.StatusNotFound, "Route not found", nil)
	})

	// Start the server
	port := ":8080"
	router.Run(port)
	logger.Log.Info("Starting server", zap.String("port", port))
}
