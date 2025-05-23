package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"go-login-api/models"
	"go-login-api/services"
)

func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	user, err := services.AuthenticateUser(&req)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			log.Println("❌ DB query error:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}
		return
	}

	accessToken, err := services.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	refreshToken := uuid.New().String()

	if err := services.DeactivatePreviousSessions(user.ID); err != nil {
		log.Println("Failed to deactivate previous sessions:", err)
	}

	// user.ID, accessToken, hashedRefresh, req.Device, req.MacAddress
	sessionID, err := services.CreateSession(models.Session{
		ID:            uuid.New().String(),
		UserID:        user.ID,
		ClientVersion: c.GetHeader("x-client-version"),
		Device:        *req.Device,
		MacAddress:    *req.MacAddress,
		PublicKey:     *req.PublicKey,
		IP:            c.RemoteIP(),
		Active:        true,
		UserAgent:     c.Request.UserAgent(),
	})
	if err != nil {
		log.Println("Insert session error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	if err := services.StoreTokenPairInRedis(user.ID, accessToken, refreshToken); err != nil {
		log.Println("Failed to store token pair in Redis:", err)
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		SessionID:    sessionID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
