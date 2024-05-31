package admin

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.POST("/", add)
		admin.GET("/", list)
		admin.GET("/:id", get)
		admin.PUT("/:id", update)
		admin.DELETE("/:id", remove)
	}
}
