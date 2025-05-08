package services

import (
	"errors"
	"hash/crc32"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"go-login-api/config"
	"go-login-api/models"
)



func AuthenticateUser(req *models.LoginRequest) (*models.User, error) {
	var user models.User

	if req.SSOID != nil {
		query := `
			SELECT id, email_hash, user_type, user_token, first_name, last_name, phone_number 
			FROM users 
			WHERE sso_id = $1`
		err := config.DB.Get(&user, query, *req.SSOID)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	if req.Email != nil && req.Password != nil {
		email := strings.ToLower(*req.Email)
		emailHash := crc32.ChecksumIEEE([]byte(email))

		query := `
			SELECT id, email_hash, password, user_type, user_token, first_name, last_name, phone_number 
			FROM users 
			WHERE email_hash = $1`
		err := config.DB.Get(&user, query, emailHash)
		if err != nil {
			return nil, err
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*req.Password)) != nil {
			return nil, bcrypt.ErrMismatchedHashAndPassword
		}
		return &user, nil
	}

	return nil, errors.New("invalid login credentials")
}

func CreateSession(userID uuid.UUID, accessToken, refreshToken string, device, macAddress *string) (string, error) {
	sessionID := uuid.New().String()

	query := `
		INSERT INTO sessions 
		(id, user_id, access_token, refresh_token, device, mac_address, active) 
		VALUES (:id, :user_id, :access_token, :refresh_token, :device, :mac_address, true)`

	_, err := config.DB.NamedExec(query, map[string]interface{}{
		"id":            sessionID,
		"user_id":       userID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"device":        device,
		"mac_address":   macAddress,
	})

	return sessionID, err
}

func DeactivatePreviousSessions(userID uuid.UUID) error {
	_, err := config.DB.Exec("UPDATE sessions SET active = false WHERE user_id = $1", userID)
	return err
}
