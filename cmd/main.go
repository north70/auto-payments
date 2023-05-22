package main

import (
	"AutoPayment/internal/repository"
	"AutoPayment/internal/repository/postgres"
	"AutoPayment/internal/repository/redis"
	"AutoPayment/internal/repository/telegram"
	"AutoPayment/internal/service"
	tgClient "AutoPayment/pkg/client/tg-client"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	redis2 "github.com/redis/go-redis/v9"
	"log"
	"os"
)

func main() {
	loadEnv()

	pgDb := loadPgDb()
	cacheDb := loadCacheDb()
	repo := repository.NewRepository(pgDb, cacheDb)
	srv := service.NewService(repo)

	loadTgBot(srv)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func loadPgDb() *sqlx.DB {
	cfg := postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	db, err := postgres.NewPostgresDb(cfg)
	if err != nil {
		panic(fmt.Sprintf("error connect to db: %s", err.Error()))
	}

	return db
}

func loadCacheDb() *redis2.Client {
	db, err := redis.NewRedisDB(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(fmt.Sprintf("error connect to cache db:%s", err.Error()))
	}

	return db
}

func loadTgBot(srv *service.Service) {
	cfg := tgClient.Config{
		Token:        os.Getenv("TELEGRAM_BOT_TOKEN"),
		DialogEnable: true,
	}

	bot := tgClient.NewBotApi(cfg)
	tg := telegram.NewTelegram(*bot, *srv)
	tg.Bot.Store = srv.Telegram
	tg.InitCommands()
	tg.HandleMessages()
}
