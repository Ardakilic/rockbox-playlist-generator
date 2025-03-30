package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	LastFMAPIKey    string
	LastFMAPISecret string
	DefaultLimit    int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		LastFMAPIKey:    os.Getenv("LASTFM_API_KEY"),
		LastFMAPISecret: os.Getenv("LASTFM_API_SECRET"),
		DefaultLimit:    50, // Default value
	}

	// Parse DEFAULT_TRACK_LIMIT if set
	if limitStr := os.Getenv("DEFAULT_TRACK_LIMIT"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DEFAULT_TRACK_LIMIT: %w", err)
		}
		config.DefaultLimit = limit
	}

	// Validate required fields
	if config.LastFMAPIKey == "" {
		return nil, fmt.Errorf("LASTFM_API_KEY is required")
	}
	if config.LastFMAPISecret == "" {
		return nil, fmt.Errorf("LASTFM_API_SECRET is required")
	}

	return config, nil
}
