package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-login-api/config"
	"go-login-api/controllers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: .env file not found, continuing with system environment variables")
	}

	// Initialize database
	if err := config.Connect(); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer config.DB.Close()

	config.InitializeRedis()

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/login", controllers.LoginHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Server error: %v", err)
	}
}
