package service

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository"
)

type TelegramService struct {
	repo repository.Telegram
}

func NewTelegramService(repo repository.Telegram) *TelegramService {
	return &TelegramService{repo: repo}
}

func (tg *TelegramService) Get(chatId int64) (model.Telegram, error) {
	return tg.repo.Get(chatId)
}

func (tg *TelegramService) UpdateAction(chatId int64, action *string) error {
	return tg.repo.UpdateAction(chatId, action)
}

func (tg *TelegramService) Upsert(chatId int64, command string, action *string) error {
	return tg.repo.Upsert(chatId, command, action)
}
