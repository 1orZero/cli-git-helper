package config

import (
	"os"
)

func LoadConfig() Config {
	return Config{
		Username:    os.Getenv("GIT_HELPER_USERNAME"),
		APIEndpoint: os.Getenv("GIT_HELPER_OPENAI_API_ENDPOINT"),
		APISecret:   os.Getenv("GIT_HELPER__OPENAI_API_SECRET"),
	}
}

type Config struct {
	Username    string
	APIEndpoint string
	APISecret   string
}

