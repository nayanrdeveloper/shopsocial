package middleware

import (
	"net/http"
	"shopsocial-backend/pkg/jwt"
	"shopsocial-backend/pkg/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware protects routes and extracts user_id from JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            responses.SendError(c, http.StatusUnauthorized, "Authorization header missing", nil)
            c.Abort()
            return
        }

        // Expecting "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            responses.SendError(c, http.StatusUnauthorized, "Invalid authorization header format", nil)
            c.Abort()
            return
        }

        tokenStr := parts[1]
        userID, err := jwt.ValidateToken(tokenStr)
        if err != nil {
            responses.SendError(c, http.StatusUnauthorized, "Invalid token", err)
            c.Abort()
            return
        }

        // Store user_id in context
        c.Set("user_id", userID)

        c.Next()
    }
}
