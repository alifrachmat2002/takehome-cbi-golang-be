package main

import (
	"log"
	"os"

	"takehome-cbi/golang-be/auth"
	"takehome-cbi/golang-be/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or error loading it: %v", err)
	}

	r := gin.Default()

	// Public route
	r.POST("/login", auth.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/items", handlers.GetItems)
		protected.POST("/items", handlers.CreateItem)
		protected.PUT("/items/:id", handlers.UpdateItem)
		protected.DELETE("/items/:id", handlers.DeleteItem)
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
