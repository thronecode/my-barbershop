package service

import (
	"github.com/thronecode/my-barbershop/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	service := router.Group("/api/service")
	service.Use(middleware.AuthMiddleware())
	{
		service.POST("/", add)
		service.GET("/", list)
		service.GET("/:id", get)
		service.PUT("/:id", update)
		service.DELETE("/:id", remove)
	}
}
