package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Postgres
	Redis
	App
}

func NewConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := Config{}
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("couldn't read config from env")
	}

	return cfg
}

type App struct {
	AppMode string `envconfig:"APP_MODE"`

	BotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
}

type Postgres struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Database string `envconfig:"DB_DATABASE"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
}

func (cfg Postgres) URI() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password)
}

type Redis struct {
	Host string `envconfig:"REDIS_HOST"`
	Port string `envconfig:"REDIS_PORT"`
}

func (cfg Redis) Address() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}
