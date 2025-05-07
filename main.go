package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"go-login-api/config"
	"go-login-api/controllers"
)

func main() {
	// connStr := "host=localhost port=5432 user=go_user password=password123 dbname=go_login sslmode=disable"
	// config.InitializeDB(connStr)
	// defer config.DB.Close()
	
	// Initialize database
	config.Connect()
	// Initialize Redis
	config.InitializeRedis()



	// Initialize Gin router
	router := gin.Default()
	router.POST("/login", controllers.LoginHandler)

	// Start the server
	log.Println("🚀 Server running at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}