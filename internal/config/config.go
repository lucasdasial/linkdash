package config

import (
	"errors"
	"os"
)

type Config struct {
	DbUrl string
}

func Load() (*Config, error) {
	dsn, hasDsn := os.LookupEnv("DATABASE_URL")

	if !hasDsn {
		return nil, errors.New("ERROR while loading config, whithout .env file")
	}

	return &Config{DbUrl: dsn}, nil
}
