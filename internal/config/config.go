package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	f := filepath.Join(homePath, ".gatorconfig.json")

	data, err := os.ReadFile(f) //([]byte, error)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	f := filepath.Join(homePath, ".gatorconfig.json")

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(f, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
