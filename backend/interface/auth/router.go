package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", Login)
	}
}
