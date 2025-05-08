package services

import (
	"context"
	"encoding/json"
	"time"

	"go-login-api/config"
)

func StoreTokenPairInRedis(userID string, accessToken, refreshToken string) error {
	ctx := context.Background()
	hashedAccessToken, _ := HashBcrypt(accessToken)
	hashedRefreshToken, _ := HashBcrypt(refreshToken)

	tokenPair := map[string]string{
		"access_token":  hashedAccessToken,
		"refresh_token": hashedRefreshToken,
	}
	jsonTokenPair, err := json.Marshal(tokenPair)
	if err != nil {
		return err
	}
	return config.RedisClient.Set(ctx, ("kunciku" + userID), jsonTokenPair, 24*time.Hour).Err()
}
