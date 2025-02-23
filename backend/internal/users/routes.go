package users

import "github.com/gin-gonic/gin"

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(router *gin.RouterGroup, handler *UserHandler) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/register", handler.RegisterUser)
        userRoutes.POST("/login", handler.LoginUser)

        // Example protected route for user details
        // userRoutes.GET("/profile", handler.GetUserProfile) // can be protected by middleware
    }
}
