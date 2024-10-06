package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	API    APIConfig    `toml:"api"`
	Branch BranchConfig `toml:"branch"`
}

type APIConfig struct {
	APIEndpoint string `toml:"api_endpoint"`
	APISecret   string `toml:"api_secret"`
}

type BranchConfig struct {
	Pattern              string `toml:"pattern"`
	DescriptionFormat    string `toml:"description_format"`
	MaxDescriptionLength int    `toml:"max_description_length"`
	NumSuggestions       int    `toml:"num_suggestions"`
}

func LoadConfig() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	configPath := filepath.Join(homeDir, ".config", "git-helper-cli", "config.toml")
	var config Config
	_, err = toml.DecodeFile(configPath, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
