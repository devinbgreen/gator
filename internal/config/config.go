package config

import "path/filepath"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	f := filepath.Join(homepath, ".gatorconfig.json")
}

func Config SetUser () {
	
}