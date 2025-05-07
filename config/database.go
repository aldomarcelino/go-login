package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)


const (
	// dbHost         = "172.16.2.52"
	// dbPort         = 5432
	// dbUser         = "app"
	// dbPassword     = "U9oe7rumHcmph2ypF0fnQwXjcmSshbEKGam9oEQsFC0BpwX45bP6EB7tEfwFpDqG"
	// dbName         = "ekuid-staging"
	// 
	dbHost         = "localhost"
	dbPort         = 5432
	dbUser         = "go_user"
	dbPassword     = "password123"
	dbName         = "go_login"
	sslMode        = "disable"
	maxAttempts    = 3
	retryDelay     = 5 * time.Second
	maxOpenConns   = 25
	maxIdleConns   = 25
	connMaxLife    = 2 * time.Hour
)

// buildConnStr constructs the PostgreSQL connection string
func buildConnStr() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
	)
}

// Connect initializes the DB connection with retry logic and pooling config
func Connect() error {
	connStr := buildConnStr()
	var err error

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
