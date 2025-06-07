package main

import (
	"Movie-Management-System/database"
	"Movie-Management-System/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	database.StartDB()

	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "Health Check API"})
	})

	router.Run(":" + port)
}
