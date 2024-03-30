package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a struct that holds the configuration for the application
type Config struct {
	AppPort string `env:"APP_PORT" env-default:"8080"`
	APIKey  string `env:"API_KEY" env-default:""`
}

// LoadConfig reads the environment variables and returns a Config struct
func LoadConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}
	return &cfg, nil
}
