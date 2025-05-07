package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)


const (
	sslMode        = "disable"
	maxAttempts    = 3
	retryDelay     = 5 * time.Second
	maxOpenConns   = 25
	maxIdleConns   = 25
	connMaxLife    = 2 * time.Hour
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

// Connect initializes the DB connection with retry logic and pooling config
func Connect() error {
	connStr, err := buildConnStr()
	if err != nil {
		return err
	}

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("[DB] Attempt %d: failed to open connection: %v", attempt, err)
			time.Sleep(retryDelay)
			continue
		}

		if err = DB.Ping(); err == nil {
			log.Println("[DB] Connection established.")
			configureConnectionPool(DB)
			return nil
		}

		log.Printf("[DB] Attempt %d: failed to ping database: %v", attempt, err)
		DB.Close()
		time.Sleep(retryDelay)
	}

	return errors.New("unable to establish DB connection after multiple attempts")
}

// configureConnectionPool sets connection pooling parameters
func configureConnectionPool(db *sql.DB) {
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLife)
}
