package config

import (
	"errors"
)

var (
	logger *Logger
)

func Init() error {
	// db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname))

	return errors.New("fake error")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
