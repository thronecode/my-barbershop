package admin

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	admin := router.Group("/api/admin")
	{
		admin.POST("/", CreateAdmin)
		admin.GET("/", ListAdmins)
		admin.GET("/:id", GetAdmin)
		admin.PUT("/:id", UpdateAdmin)
		admin.DELETE("/:id", DeleteAdmin)
	}
}
