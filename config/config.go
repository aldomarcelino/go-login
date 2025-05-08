package config

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB          *sqlx.DB
	RedisClient *redis.Client
	JWTSecret   = []byte("aldobareto01")
	JWTIssuer   = "go-login-api"
)

func InitializeDB(connStr string) {
	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}


func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}