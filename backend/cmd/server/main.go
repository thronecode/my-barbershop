package main

import (
	"github.com/thronecode/my-barbershop/backend/internal/config"
	"github.com/thronecode/my-barbershop/backend/internal/config/database"
	_ "github.com/thronecode/my-barbershop/backend/internal/docs"
	"github.com/thronecode/my-barbershop/backend/internal/interface/admin"
	"github.com/thronecode/my-barbershop/backend/internal/interface/auth"
	"github.com/thronecode/my-barbershop/backend/internal/interface/barber"
	"github.com/thronecode/my-barbershop/backend/internal/interface/service"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My BarberShop API
// @version 1.0
// @description API for a My BarberShop application

// @host localhost:4002
// @BasePath /api

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := database.OpenConnections(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.CloseConnections()

	router := gin.Default()
	admin.RegisterRoutes(router)
	auth.RegisterRoutes(router)
	barber.RegisterRoutes(router)
	service.RegisterRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":4002"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}