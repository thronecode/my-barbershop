package main

import (
	"backend/config"
	"backend/config/database"
	_ "backend/docs"
	"backend/interface/admin"
	"backend/interface/auth"
	"backend/interface/barber"

	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My BarberShop API
// @version 1.0
// @description API for a My BarberShop application

// @host localhost:4002
// @BasePath /api

func main() {
	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := database.OpenConnections(); err != nil {
		//log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.CloseConnections()

	router := gin.Default()
	admin.RegisterRoutes(router)
	auth.RegisterRoutes(router)
	barber.RegisterRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":4002"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
