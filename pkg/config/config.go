package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Token   string
	GroupID string
}

// GetConfig reads configs from env variables
func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error reading env vars")
	}
	token := os.Getenv("VK_BOT_TOKEN")
	groupID := os.Getenv("VK_BOT_GROUP_ID")

	if token == "" {
		return nil, errors.New("no token set")
	}
	if groupID == "" {
		return nil, errors.New("no group id set")
	}

	return &Config{
		Token:   token,
		GroupID: groupID,
	}, nil
}
