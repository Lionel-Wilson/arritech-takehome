package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT" yaml:"port" validate:"required"`
	DatabaseURL string `mapstructure:"DATABASE_URL" yaml:"database_url" validate:"required"`
	LogLevel    string `mapstructure:"LOG_LEVEL" yaml:"log_level"`
	Env         string `mapstructure:"ENV" yaml:"env" validate:"required"`
}

// LoadConfig loads from OS env; if ENV=local (or unset) it will attempt to load .env first.
func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	// Sensible default; can be overridden by real ENV
	viper.SetDefault("ENV", "local")

	// If ENV explicitly set to "local" (or not set in OS), try .env without failing hard.
	rawEnv := os.Getenv("ENV")
	if rawEnv == "" || strings.EqualFold(rawEnv, "local") {
		if err := godotenv.Load(); err != nil {
			// Not fatal—just informational in local dev
			log.Printf("no .env file found (ok if running in CI/containers): %v", err)
		}
	}

	// Create a Config instance with values from environment variables.
	cfg := Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		LogLevel:    viper.GetString("LOG_LEVEL"),
		Env:         viper.GetString("ENV"),
	}

	// Validate the config.
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}
