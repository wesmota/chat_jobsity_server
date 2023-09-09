package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Postgres driver for initializing database connection
	_ "github.com/lib/pq"
)

const (
	appName         = "go-jobsity-chat-server"
	localDBUser     = "root"
	localDBPassword = "password"
	localDBHost     = "localhost"
	localPort       = "5432"
	localDBName     = "jobsity"
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

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = localPort
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = localDBName
	}

	dbConnectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)
	log.Println("Connecting to database with connection string: ", dbConnectionStr)

	return sql.Open("postgres", dbConnectionStr)
}
