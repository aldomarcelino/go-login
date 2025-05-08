package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"go-login-api/config"
	"go-login-api/models"
)

func CryptoHash(inputString string) string {
	hasher := sha512.New()
	hasher.Write([]byte(inputString))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HashBcrypt hashes a password using bcrypt
func HashBcrypt(password string) (string, error) {
	// Generate salt and hash in one step (bcrypt.GenerateFromPassword does both)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Wrap error into InternalServerError
		return "", err
	}

	return string(hashedPassword), nil
}

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
		emailHash := CryptoHash(email)

		query := `
			SELECT id, email_hash, password, first_name, last_name, phone_number 
			FROM "Users" 
			WHERE email_hash = $1`
		// query := `SELECT * FROM users limit 1`
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

func CreateSession(session models.Session) (string, error) {
	sessionID := uuid.New().String()

	// query := `
	// 	INSERT INTO "Sessions"
	// 	(id, user_id, access_token, refresh_token, device, mac_address, active)
	// 	VALUES (:id, :user_id, :access_token, :refresh_token, :device, :mac_address, true)`

	query := `
		INSERT INTO "Sessions"
		(id, user_id, client_version, device, mac_address, public_key, active, ip, user_agent)
		VALUES (:id, :user_id, :client_version, :device, :mac_address, :public_key, :active, :ip, :user_agent)
	`
	_, err := config.DB.NamedExec(query, session)

	return sessionID, err
}

func DeactivatePreviousSessions(userID string) error {
	_, err := config.DB.Exec(`UPDATE "Sessions" SET active = false WHERE user_id = $1`, userID)
	return err
}
