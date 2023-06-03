package telegram

import (
	"AutoPayment/internal/handler/telegram/command"
	"AutoPayment/internal/service"
	"AutoPayment/pkg/client/tg-client"
)

type Telegram struct {
	Bot     tg_client.BotApi
	Service service.Service
}

func NewTelegram(bot tg_client.BotApi, serv service.Service) *Telegram {
	return &Telegram{Bot: bot, Service: serv}
}

func (tg *Telegram) InitCommands() {
	tg.Bot.Commands.Commands = make(map[string]tg_client.Command)

	tg.Bot.Commands.AddMany([]tg_client.Command{
		command.NewHelpCommand(tg.Bot),
		command.NewPaymentNewCommand(tg.Bot, tg.Service.PaymentTemp, tg.Service.Payment),
		command.NewListPaymentCommand(tg.Bot, tg.Service),
		command.NewWhoamiCommand(tg.Bot),
	})
}

func (tg *Telegram) HandleMessages() {
	query := tg_client.UpdateQuery{
		Offset:  0,
		Limit:   10,
		Timeout: 5,
	}

	updates := tg.Bot.GetUpdatesChan(query)

	for update := range updates {
		err := tg.Bot.HandleMessage(update)
		if err != nil {
			tg.Bot.Logger.Error().Msg(err.Error())
		}
	}
}
