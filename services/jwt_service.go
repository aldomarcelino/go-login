package services

import (
	"time"

	"go-login-api/config"
	"go-login-api/models"
	"go-login-api/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateAccessToken(user *models.User) (string, error) {
	decryptedFirstName := utils.DecryptMock(user.FirstName.String)
	decryptedLastName := utils.DecryptMock(user.LastName.String)
	decryptedPhoneNumber := utils.DecryptMock(user.PhoneNumber.String)

	tokenPayload := jwt.MapClaims{
		"user_id":      user.ID.String(),
		"user_type":    user.UserType.String,
		"user_token":   user.UserToken.String,
		"first_name":   decryptedFirstName,
		"last_name":    decryptedLastName,
		"phone_number": decryptedPhoneNumber,
		"exp":          time.Now().Add(15 * time.Minute).Unix(),
		"iss":          config.JWTIssuer,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenPayload)
	return accessToken.SignedString(config.JWTSecret)
}

func GenerateRefreshToken() (string, string, error) {
	rawRefresh := uuid.New().String()
	hashedRefresh, err := bcrypt.GenerateFromPassword([]byte(rawRefresh), bcrypt.DefaultCost)
	return rawRefresh, string(hashedRefresh), err
}