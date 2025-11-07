package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds application configuration values.
type Config struct {
	Port           string
	DBPath         string
	AllowedOrigins []string
	BaseURL        string
}

// Load reads configuration from environment variables and optional .env file.
func Load() (*Config, error) {
	_ = godotenv.Load()

	port := getenv("PORT", "8080")
	dbPath := getenv("DB_PATH", "./urls.db")
	corsOrigins := getenv("CORS_ORIGINS", "http://localhost:3001")
	baseURL := getenv("BASE_URL", "http://localhost:8080")

	cfg := &Config{
		Port:           port,
		DBPath:         dbPath,
		AllowedOrigins: splitAndTrim(corsOrigins),
		BaseURL:        baseURL,
	}

	return cfg, nil
}

// Addr returns the address the server should listen on.
func (c *Config) Addr() string {
	return fmt.Sprintf(":%s", c.Port)
}

func getenv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}

func splitAndTrim(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
