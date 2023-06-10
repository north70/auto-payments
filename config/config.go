package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres
	Redis
	App
}

func NewConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, errors.New("error loading .env file")
	}

	cfg := Config{}
	err = envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, errors.New("couldn't read config from env")
	}

	return cfg, nil
}
