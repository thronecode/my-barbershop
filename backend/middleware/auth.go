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
			sorry.Handling(c, sorry.NewErr("Authorization header required", http.StatusUnauthorized))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !utils.IsValidToken(token) {
			sorry.Handling(c, sorry.NewErr("Invalid token", http.StatusUnauthorized))
			return
		}

		c.Next()
	}
}
