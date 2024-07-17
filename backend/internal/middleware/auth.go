package middleware

import (
	"github.com/thronecode/my-barbershop/backend/internal/sorry"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the authentication token
func AuthMiddleware(validateToken func(string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			sorry.Handling(c, sorry.NewErr("Authorization header required", http.StatusUnauthorized))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !validateToken(token) {
			sorry.Handling(c, sorry.NewErr("Invalid token", http.StatusUnauthorized))
			return
		}

		c.Next()
	}
}
