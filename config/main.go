package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type DbConfig struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     string
}

type ValkeyConfig struct {
	VkAddress  string
	VkPassword string
}

type EnvConfig struct {
	ExternalAPISecret string
	DbConfig
	ValkeyConfig
}

// ENVIRONMENT VARIABLE NAMES
const RAWG_API_SECRET = "RAWG_API_SECRET"
const DB_HOST = "DB_HOST"
const DB_PORT = "DB_PORT"
const DB_NAME = "DB_NAME"
const DB_USER = "DB_USER"
const DB_PASSWORD = "DB_PASSWORD"

const VALKEY_HOST = "VALKEY_HOST"
const VALKEY_PORT = "VALKEY_PORT"
const VALKEY_PASSWORD = "VALKEY_PASSWORD"

const JWT_SECRET = "JWT_SECRET"

// LoadEnvConfig returns all the environment variables in a struct format
func LoadEnvConfig() (*EnvConfig, error) {

	cfg := &EnvConfig{
		ExternalAPISecret: os.Getenv(RAWG_API_SECRET),
		DbConfig: DbConfig{
			DbHost:     os.Getenv(DB_HOST),
			DbPort:     os.Getenv(DB_PORT),
			DbName:     os.Getenv(DB_NAME),
			DbUser:     os.Getenv(DB_USER),
			DbPassword: os.Getenv(DB_PASSWORD),
		},
		ValkeyConfig: ValkeyConfig{
			VkAddress:  fmt.Sprintf("%s:%s", os.Getenv(VALKEY_HOST), os.Getenv(VALKEY_PORT)),
			VkPassword: os.Getenv(VALKEY_PASSWORD),
		},
	}

	checks := []bool{
		cfg.ExternalAPISecret == "",
		cfg.DbConfig.DbHost == "",
		cfg.DbConfig.DbPort == "",
		cfg.DbConfig.DbName == "",
		cfg.DbConfig.DbUser == "",
		cfg.DbConfig.DbPassword == "",
		strings.HasPrefix(cfg.ValkeyConfig.VkAddress, ":"),
		strings.HasSuffix(cfg.ValkeyConfig.VkAddress, ":"),
	}

	for _, check := range checks {
		if check {
			return nil, errors.New("Unable to get required environment variables")
		}
	}

	return cfg, nil

}
