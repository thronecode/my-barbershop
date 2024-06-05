package barber

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	barber := router.Group("/api/barber")
	barber.Use(middleware.AuthMiddleware())
	{
		barber.POST("/", add)
		barber.GET("/", list)
		barber.GET("/:id", get)
		barber.PUT("/:id", update)
		barber.DELETE("/:id", remove)
		barber.POST("/:id/checkin", addCheckin)
		barber.GET("/:id/checkin", getCheckins)
		barber.POST("/:id/service", addService)
		barber.DELETE("/:id/service/:service_id", removeService)
	}
}
