package utils

import (
	"errors"
	"os"
)

type Config struct {
	ExternalAPISecret string
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		ExternalAPISecret: os.Getenv("RAWG_API_SECRET"),
	}

	if cfg.ExternalAPISecret == "" {
		return nil, errors.New("Unable to get required environment variables")
	}

	return cfg, nil

}
