package main

import (
	"AutoPayment/config"
	"AutoPayment/internal/handler/telegram"
	"AutoPayment/internal/repository"
	"AutoPayment/internal/repository/postgres"
	"AutoPayment/internal/repository/redis"
	"AutoPayment/internal/service"
	tgClient "AutoPayment/pkg/client/tg-client"
	"AutoPayment/pkg/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	redis2 "github.com/redis/go-redis/v9"
)

func main() {
	cfg := config.NewConfig()

	pgDb := loadPgDb(cfg.Postgres)
	cacheDb := loadCacheDb(cfg.Redis)
	repo := repository.NewRepository(pgDb, cacheDb)
	srv := service.NewService(repo)

	loadTgBot(srv, cfg.App)
}

func loadPgDb(cfg config.Postgres) *sqlx.DB {
	db, err := postgres.NewPostgresDb(cfg)
	if err != nil {
		panic(fmt.Sprintf("error connect to db: %s", err.Error()))
	}

	return db
}

func loadCacheDb(cfg config.Redis) *redis2.Client {
	db, err := redis.NewRedisDB(cfg)
	if err != nil {
		panic(fmt.Sprintf("error connect to cache db:%s", err.Error()))
	}

	return db
}

func loadTgBot(srv *service.Service, cfgApp config.App) {
	l := logger.InitLogger()

	cfgTg := tgClient.Config{
		Token:        cfgApp.BotToken,
		DialogEnable: true,
	}

	bot := tgClient.NewBotApi(cfgTg)
	bot.Logger = l
	tg := telegram.NewTelegram(*bot, *srv)
	tg.Bot.Store = srv.Telegram
	tg.InitCommands()
	tg.HandleMessages()
}
