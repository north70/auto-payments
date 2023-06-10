package command

import (
	"AutoPayment/internal/handler/telegram/action"
	"AutoPayment/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command interface {
	Name() string
	Description() string
	Handle(upd tgbotapi.Update) error
	NextAction() *string
}

type BaseCommand struct {
	TGBot   *tgbotapi.BotAPI
	Service *service.Service
	Actions map[string]action.Action
}

func NewBaseCommand(TGBot *tgbotapi.BotAPI, service *service.Service, actions map[string]action.Action) *BaseCommand {
	return &BaseCommand{TGBot: TGBot, Service: service, Actions: actions}
}

func (c *BaseCommand) Name() string {
	return ""
}

func (c *BaseCommand) Description() string {
	return ""
}

func (c *BaseCommand) Handle(upd tgbotapi.Update) error {
	return nil
}

func (c *BaseCommand) NextAction() *string {
	return nil
}
