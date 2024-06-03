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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			sorry.Handling(c, sorry.Err(&sorry.Error{
				Code:       sorry.ValidationErroCode,
				Err:        sorry.NewErr("Authorization header required"),
				Msg:        "Authorization header required",
				StatusCode: http.StatusUnauthorized,
			}))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !utils.IsValidToken(token) {
			sorry.Handling(c, sorry.Err(&sorry.Error{
				Code:       sorry.ValidationErroCode,
				Err:        sorry.NewErr("Invalid token"),
				Msg:        "Invalid token",
				StatusCode: http.StatusUnauthorized,
			}))
			return
		}

		c.Next()
	}
}
