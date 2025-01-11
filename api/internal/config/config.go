package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl        string
	PasetoSecret string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		DbUrl:        os.Getenv("DB_URL"),
		PasetoSecret: os.Getenv("PASETO_SECRET_KEY"),
	}, nil
}
