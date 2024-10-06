package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type DatabaseConfig struct {
	DbName   string
	Port     int
	Host     string
	User     string
	Password string
}

type Config struct {
	Env      string
	LogLevel string
	Database DatabaseConfig
}

func getStrEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseInt(value, 10, 16); err == nil {
			return int(result)
		}
	}
	return fallback
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		Env:      getStrEnv("ENV", "local"),
		LogLevel: getStrEnv("LOG_LEVEL", "debug"),
		Database: DatabaseConfig{
			DbName:   getStrEnv("DB_NAME", "clean_architecture"),
			Port:     getIntEnv("DB_PORT", 5432),
			Host:     getStrEnv("DB_HOST", "localhost"),
			User:     getStrEnv("DB_USER", "postgres"),
			Password: getStrEnv("DB_PASSWORD", "postgres"),
		},
	}
}

var ConfigModule = fx.Module("ConfigModule", fx.Option(fx.Provide(NewConfig)))
