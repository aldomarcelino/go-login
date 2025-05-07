package config

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var (
	DB          *sql.DB
	RedisClient *redis.Client
	JWTSecret   = []byte("aldobareto01")
	JWTIssuer   = "go-login-api"
)

func InitializeDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}