package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIKey string `yaml:"api_key"`
}

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "mycli", "config.yaml")
}

func GetAPIKey() string {
	path := getConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	var cfg Config
	yaml.Unmarshal(data, &cfg)
	return cfg.APIKey
}
