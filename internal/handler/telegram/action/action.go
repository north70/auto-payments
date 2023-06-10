package action

import (
	"AutoPayment/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Action interface {
	Name() string
	Handle(upd tgbotapi.Update) error
	Next() string
}

type BaseAction struct {
	TGBot   *tgbotapi.BotAPI
	Service *service.Service
}

func NewBaseAction(TGBot *tgbotapi.BotAPI, service *service.Service) *BaseAction {
	return &BaseAction{TGBot: TGBot, Service: service}
}

func (a *BaseAction) Name() string {
	return ""
}

func (a *BaseAction) Handle(upd tgbotapi.Update) error {
	return nil
}

func (a *BaseAction) Next() string {
	return ""
}
