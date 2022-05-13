package configs

import (
	"fmt"
	"os"
)

type Config struct {
	// Database access URL
	DBUrl string

	// Server port
	ListenAddr string
}

// GetConfig gets the app configuration.
func GetConfig() *Config {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	endpoint := os.Getenv("DB_ENDPOINT")
	initDB := os.Getenv("DB_INIT")

	url := fmt.Sprintf("postgres://%s:%s@%s/%s",
		user, password, endpoint, initDB)

	addr := ":5000"

	return &Config{
		DBUrl:      url,
		ListenAddr: addr,
	}
}
