package middleware

import (
	"backend/sorry"
	"backend/utils"

	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
)

// AuthMiddleware verifies the authentication token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": sorry.NewErr("Authorization header required")})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !utils.IsValidToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": sorry.NewErr("Invalid token")})
			c.Abort()
			return
		}

		c.Next()
	}
}
