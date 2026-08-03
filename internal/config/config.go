package config

import (
	"errors"
	"os"
)

type Config struct {
	DbUrl string
	Port  string
}

func Load() (*Config, error) {

	dsn, hasDsn := os.LookupEnv("DATABASE_URL")
	port := "8080"

	if !hasDsn {
		return nil, errors.New("ERROR while loading config, whithout .env file")
	}

	if portEnv, hasPort := os.LookupEnv("PORT"); hasPort {
		port = portEnv
	}

	return &Config{DbUrl: dsn, Port: port}, nil
}
