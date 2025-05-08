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
	firstName := ""
	lastName := ""
	phoneNumber := ""
	userType := ""
	userToken := ""

	if user.FirstName!=nil {
		firstName = utils.DecryptMock(*user.FirstName)
	}
	if user.LastName!=nil {
		lastName = utils.DecryptMock(*user.LastName)
	}
	if user.PhoneNumber!=nil {
		phoneNumber = utils.DecryptMock(*user.PhoneNumber)
	}
	if user.UserType!=nil {
		userType = *user.UserType
	}
	if user.UserToken!=nil {
		userToken = *user.UserToken
	}

	tokenPayload := jwt.MapClaims{
		"user_id":      user.ID.String(),
		"user_type":    userType,
		"user_token":   userToken,
		"first_name":   firstName,
		"last_name":    lastName,
		"phone_number": phoneNumber,
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