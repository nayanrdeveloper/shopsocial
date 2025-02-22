package users

import "github.com/gin-gonic/gin"

// RegisterUserRoutes registers routes for the user module
func RegisterUserRoutes(router *gin.RouterGroup, handler *UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/:email", handler.GetUserByEmail)
	}
}
