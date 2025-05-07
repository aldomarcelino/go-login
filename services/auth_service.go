package services

import (
	"database/sql"
	"hash/crc32"
	"log"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"go-login-api/config"
	"go-login-api/models"
)

func AuthenticateUser(req *models.LoginRequest) (*models.User, error) {
	var user models.User
	var err error

	if req.SSOID != nil {
		err = config.DB.QueryRow("SELECT id, email_hash, user_type, user_token, first_name, last_name, phone_number FROM users WHERE sso_id = $1", *req.SSOID).
			Scan(&user.ID, &user.EmailHash, &user.UserType, &user.UserToken, &user.FirstName, &user.LastName, &user.PhoneNumber)
	} else if req.Email != nil && req.Password != nil {
		email := strings.ToLower(*req.Email)
		emailHash := crc32.ChecksumIEEE([]byte(email))

		log.Println("Email:", *req.Email)
		log.Println("Email Hash:", emailHash)

		err = config.DB.QueryRow(`SELECT id, email_hash, password, user_type, user_token, first_name, last_name, phone_number 
		                   FROM users WHERE email_hash = $1`, emailHash).
			Scan(&user.ID, &user.EmailHash, &user.Password, &user.UserType, &user.UserToken, &user.FirstName, &user.LastName, &user.PhoneNumber)

		if err == nil && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*req.Password)) != nil {
			return nil, bcrypt.ErrMismatchedHashAndPassword
		}
	} else {
		return nil, sql.ErrNoRows
	}

	return &user, err
}

func CreateSession(userID uuid.UUID, accessToken, refreshToken string, device, macAddress *string) (string, error) {
	sessionID := uuid.New().String()
	_, err := config.DB.Exec(`INSERT INTO sessions 
	                  (id, user_id, access_token, refresh_token, device, mac_address, active) 
	                  VALUES ($1, $2, $3, $4, $5, $6, true)`,
		sessionID, userID, accessToken, refreshToken, device, macAddress)
	return sessionID, err
}

func DeactivatePreviousSessions(userID uuid.UUID) error {
	_, err := config.DB.Exec("UPDATE sessions SET active = false WHERE user_id = $1", userID)
	return err
}