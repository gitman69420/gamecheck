package utils

import (
	"errors"
	"os"
)

type EnvConfig struct {
	ExternalAPISecret string
}

func LoadEnvConfig() (*EnvConfig, error) {

	cfg := &EnvConfig{
		ExternalAPISecret: os.Getenv("RAWG_API_SECRET"),
	}

	if cfg.ExternalAPISecret == "" {
		return nil, errors.New("Unable to get required environment variables")
	}

	return cfg, nil

}
