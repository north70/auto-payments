package command

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/service"
	tg_client "AutoPayment/pkg/client/tg-client"
	"fmt"
)

type ListPaymentCommand struct {
	Bot     tg_client.BotApi
	Service service.Service
}

func NewListPaymentCommand(bot tg_client.BotApi, service service.Service) *ListPaymentCommand {
	return &ListPaymentCommand{Bot: bot, Service: service}
}

func (cmd *ListPaymentCommand) Name() string {
	return "list"
}

func (cmd *ListPaymentCommand) Description() string {
	return "Список всех платежей"
}

func (cmd *ListPaymentCommand) Handle(update tg_client.Update) error {
	userId := update.Message.From.Id
	payments, err := cmd.Service.Payment.Index(userId)

	if err != nil || len(payments) == 0 {
		sendMsgQuery := tg_client.SendMessageQuery{
			ChatId: update.Message.From.Id,
			Text:   "Не найдено платежей",
		}

		err = cmd.Bot.SendMessage(sendMsgQuery)
		if err != nil {
			return err
		}
	}

	message := formatPayments(payments)

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: update.Message.From.Id,
		Text:   message,
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func formatPayments(payments []model.Payment) string {
	var message string

	for _, payment := range payments {
		message = fmt.Sprintf("%s \n"+payment.StringForTg(), message)
	}

	return message
}
