package main

import (
	"backend/config"
	"backend/config/database"
	"backend/interface/admin"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := database.OpenConnections(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.CloseConnections()

	router := gin.Default()
	admin.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
