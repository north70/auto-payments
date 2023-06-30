package main

import (
	"AutoPayment/config"
	"AutoPayment/internal/handler/scheduler"
	"AutoPayment/internal/handler/telegram"
	"AutoPayment/internal/repository"
	"AutoPayment/internal/service"
	"AutoPayment/internal/storage"
	"AutoPayment/pkg/logger"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func main() {
	log := logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("config loaded")

	pgDb, err := storage.NewPostgresDb(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("error connect to db")
	}
	log.Info().Msg("postgres database connected")

	cacheDb, err := storage.NewRedisDB(cfg.Redis)
	if err != nil {
		log.Fatal().Err(err).Msg("error connect to cache db:%s")
	}
	log.Info().Msg("redis database connected")

	repo := repository.NewRepository(pgDb, cacheDb)
	srv := service.NewService(repo)

	location, err := time.LoadLocation(cfg.App.AppLocation)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("error load location %s", cfg.App.AppLocation))
	}
	log.Info().Msg(fmt.Sprintf("loaded location %s", location.String()))

	schedule := scheduler.NewScheduler(srv, log, location)
	schedule.Start()
	log.Info().Msg("scheduler started")

	botApi, err := tgbotapi.NewBotAPI(cfg.App.BotToken)
	if err != nil {
		log.Fatal().Msg("error authorized in telegram")
	}
	log.Info().Msg("telegram success authorized")

	bot := telegram.NewTgBot(botApi, cfg, log, srv)
	bot.Start()
}
