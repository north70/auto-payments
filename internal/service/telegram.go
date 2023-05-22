package service

import (
	"AutoPayment/internal/repository"
	tg_client "AutoPayment/pkg/client/tg-client"
)

type TelegramService struct {
	repo repository.Telegram
}

func NewTelegramService(repo repository.Telegram) *TelegramService {
	return &TelegramService{repo: repo}
}

func (tg *TelegramService) Has(chatId int) (bool, error) {
	return tg.repo.Has(chatId)
}

func (tg *TelegramService) New(chatId int, command string) error {
	return tg.repo.Create(chatId, command)
}

func (tg *TelegramService) Current(chatId int) (tg_client.ChatStatus, error) {
	tgStatus, err := tg.repo.Get(chatId)

	if err != nil {
		return tg_client.ChatStatus{}, err
	}

	return tg_client.ChatStatus{
		Command: tgStatus.Command,
		Action:  tgStatus.Action,
	}, nil
}

func (tg *TelegramService) ClearAction(chatId int) error {
	return tg.repo.ClearAction(chatId)
}

func (tg *TelegramService) SetCommand(chatId int, command string) error {
	return tg.repo.Update(chatId, &command, nil)
}

func (tg *TelegramService) SetAction(chatId int, action string) error {
	return tg.repo.Update(chatId, nil, &action)
}
