package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	appName         = "go-jobsity-chat-server"
	localDBUser     = "root"
	localDBPassword = "password"
	localDBHost     = "localhost"
)

// NewDB returns an initialized database handle
func NewDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = localDBHost
	}
	return initDB(dbHost)
}

func initDB(host string) (*sql.DB, error) {
	var dbUser, dbPassword string

	dbUser = os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = localDBUser
	}

	dbPassword = os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = localDBPassword
	}

	dbConnectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		os.Getenv("DB_PORT"),
		dbUser,
		dbPassword,
		os.Getenv("DB_NAME"),
	)

	return sql.Open("postgres", dbConnectionStr)
}
