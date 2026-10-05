package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultHTTPHost = ""
	defaultHTTPPort = 8080
)

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Host string
	Port int
}

// Load reads and validates the application configuration from environment
// variables, applying defaults when optional values are not provided.
func Load() (Config, error) {
	host := getEnv("HTTP_HOST", defaultHTTPHost)

	port, err := getIntEnv("HTTP_PORT", defaultHTTPPort)
	if err != nil {
		return Config{}, err
	}

	if port < 1 || port > 65535 {
		return Config{}, fmt.Errorf(
			"HTTP_PORT must be between 1 and 65535, got %d",
			port,
		)
	}

	return Config{
		HTTP: HTTPConfig{
			Host: host,
			Port: port,
		},
	}, nil
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	return value
}

func getIntEnv(key string, fallback int) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer: %w", key, err)
	}

	return parsed, nil
}
