package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	logger *Logger
)

var (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "password" // Replace with your actual password
	dbname   string = "mydatabase"
)

func Init() error {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname))

	if err != nil {
		logger.Errorf("could not start Postgres: %v", err)
	}

	logger.Info(db)

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
