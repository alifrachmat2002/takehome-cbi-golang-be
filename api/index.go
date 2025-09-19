package handler

import (
	"log"
	"net/http"

	"takehome-cbi/golang-be/auth"
	"takehome-cbi/golang-be/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	// Load environment variables (optional in production)
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configure CORS for production
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Public route
	router.POST("/login", auth.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/items", handlers.GetItems)
		protected.GET("/item/:id", handlers.GetItem)
		protected.POST("/items", handlers.CreateItem)
		protected.PUT("/items/:id", handlers.UpdateItem)
		protected.DELETE("/items/:id", handlers.DeleteItem)
	}

	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Takehome CBI Golang Backend API",
			"status":  "healthy",
		})
	})
}

// Handler adalah entry point untuk Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}