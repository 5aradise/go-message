package config

import (
	"errors"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	DB struct {
		Path string
	}
}

func LoadFromEnv() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return Config{}, errors.New("PORT is not found in the enviroment")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		return Config{}, errors.New("DB_PATH is not found in the enviroment")
	}

	cfg := Config{}

	cfg.Server.Port = port
	cfg.DB.Path = dbPath

	return cfg, nil
}
