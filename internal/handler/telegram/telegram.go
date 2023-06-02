package telegram

import (
	command2 "AutoPayment/internal/handler/telegram/command"
	"AutoPayment/internal/service"
	"AutoPayment/pkg/client/tg-client"
	"fmt"
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
		command2.NewHelpCommand(tg.Bot),
		command2.NewPaymentNewCommand(tg.Bot),
		command2.NewListPaymentCommand(tg.Bot, tg.Service),
		command2.NewWhoamiCommand(tg.Bot),
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
			fmt.Println(err.Error())
		}
	}
}
