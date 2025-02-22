package main

import (
	"shopsocial-backend/api"
	"shopsocial-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to MongoDB
	config.ConnectDB()

	// Setup Gin
	router := gin.Default()

	// Register all routes
	api.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
