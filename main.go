package main

import (
	"Movie-Management-System/database"
	"Movie-Management-System/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
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
