package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Port string
	}
	DB struct {
		Path string
	}
	JWT struct {
		Key []byte
	}
	Auth struct {
		AccessTokenMaxAge  int
		RefreshTokenMaxAge int
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

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return Config{}, errors.New("JWT_SECRET is not found in the enviroment")
	}

	accMaxAgeStr := os.Getenv("ACCESS_TOKEN_MAX_AGE")
	if accMaxAgeStr == "" {
		return Config{}, errors.New("ACCESS_TOKEN_MAX_AGE is not found in the enviroment")
	}
	accMaxAge, err := strconv.Atoi(accMaxAgeStr)
	if err != nil {
		return Config{}, err
	}

	refMaxAgeStr := os.Getenv("REFRESH_TOKEN_MAX_AGE")
	if refMaxAgeStr == "" {
		return Config{}, errors.New("REFRESH_TOKEN_MAX_AGE is not found in the enviroment")
	}
	refMaxAge, err := strconv.Atoi(refMaxAgeStr)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	cfg.Server.Port = port
	cfg.DB.Path = dbPath
	cfg.JWT.Key = []byte(jwtKey)
	cfg.Auth.AccessTokenMaxAge = accMaxAge
	cfg.Auth.RefreshTokenMaxAge = refMaxAge

	return cfg, nil
}
