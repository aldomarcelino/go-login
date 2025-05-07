package services

import (
	"context"
	"encoding/json"
	"time"

	"go-login-api/config"
)

func StoreTokenPairInRedis(userID string, accessToken, refreshToken string) error {
	ctx := context.Background()
	tokenPair := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	jsonTokenPair, err := json.Marshal(tokenPair)
	if err != nil {
		return err
	}
	return config.RedisClient.Set(ctx, ("kunciku"+userID), jsonTokenPair, 24*time.Hour).Err()
}