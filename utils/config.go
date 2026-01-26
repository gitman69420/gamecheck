package utils

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

func LoadEnvConfig() (*EnvConfig, error) {

	cfg := &EnvConfig{
		ExternalAPISecret: os.Getenv("RAWG_API_SECRET"),
		DbConfig: DbConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
		ValkeyConfig: ValkeyConfig{
			VkAddress:  fmt.Sprintf("%s:%s", os.Getenv("VALKEY_HOST"), os.Getenv("VALKEY_PORT")),
			VkPassword: os.Getenv("VALKEY_PASSWORD"),
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
