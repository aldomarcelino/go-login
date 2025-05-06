package main

import (
	"go-login/config"
	"go-login/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	r.Run(":8080")
}
