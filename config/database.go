package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	sslMode      = "disable"
	maxAttempts  = 3
	retryDelay   = 5 * time.Second
	maxOpenConns = 25
	maxIdleConns = 25
	connMaxLife  = 2 * time.Hour
)


// buildConnStr constructs the PostgreSQL connection string from environment variables
func buildConnStr() (string, error) {
	dbHost := os.Getenv("DB_STAG_HOST")
	dbPort := os.Getenv("DB_STAG_PORT")
	dbUser := os.Getenv("DB_STAG_USER")
	dbPassword := os.Getenv("DB_STAG_PASS")
	dbName := os.Getenv("DB_STAG_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return "", errors.New("missing required environment variables for database connection")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
	), nil
}

// Connect initializes the DB connection using sqlx with retry and pooling configuration
func Connect() error {
	connStr, err := buildConnStr()
	if err != nil {
		return err
	}

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		var db *sqlx.DB
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			log.Println("[DB] Connection established.")
			configureConnectionPool(db)
			DB = db
			return nil
		}

		log.Printf("[DB] Attempt %d: failed to connect: %v", attempt, err)
		time.Sleep(retryDelay)
	}

	return errors.New("unable to establish DB connection after multiple attempts")
}

// configureConnectionPool sets connection pooling parameters
func configureConnectionPool(db *sqlx.DB) {
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLife)
}
