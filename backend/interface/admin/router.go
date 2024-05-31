package admin

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	admin := router.Group("/api/admin")
	{
		admin.POST("/", add)
		admin.GET("/", list)
		admin.GET("/:id", get)
		admin.PUT("/:id", update)
		admin.DELETE("/:id", remove)
	}
}
