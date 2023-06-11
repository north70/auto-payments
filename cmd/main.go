package main

import (
	"AutoPayment/config"
	"AutoPayment/internal/handler/telegram"
	"AutoPayment/internal/repository"
	"AutoPayment/internal/service"
	"AutoPayment/internal/storage"
	"AutoPayment/pkg/logger"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	redis2 "github.com/redis/go-redis/v9"
)

func main() {
	log := logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("config loaded")

	pgDb, err := loadPgDb(cfg.Postgres)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("postgres database connected")

	cacheDb, err := loadCacheDb(cfg.Redis)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("redis database connected")

	repo := repository.NewRepository(pgDb, cacheDb)
	srv := service.NewService(repo)

	botApi, err := tgbotapi.NewBotAPI(cfg.App.BotToken)
	if err != nil {
		log.Fatal().Msg("error authorized in telegram")
	}
	log.Info().Msg("telegram success authorized")

	bot := telegram.NewTgBot(botApi, cfg, log, srv)
	bot.Start()
}

func loadPgDb(cfg config.Postgres) (*sqlx.DB, error) {
	db, err := storage.NewPostgresDb(cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error connect to db: %s", err.Error()))
	}

	return db, nil
}

func loadCacheDb(cfg config.Redis) (*redis2.Client, error) {
	db, err := storage.NewRedisDB(cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error connect to cache db:%s", err.Error()))
	}

	return db, nil
}
