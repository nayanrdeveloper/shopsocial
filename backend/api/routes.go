package api

import (
	"shopsocial-backend/internal/products"
	"shopsocial-backend/internal/users"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API routes for the application
func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api") // Prefix for all API routes

	// Initialize Users Module
	userRepo := users.NewUserRepository()
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)
	users.RegisterUserRoutes(apiGroup, userHandler)

	// Initialize Products Module
	productRepo := products.NewProductRepository()
	productService := products.NewProductService(productRepo)
	productHandler := products.NewProductHandler(productService)
	products.RegisterProductRoutes(apiGroup, productHandler)
}
