package command

import (
	"AutoPayment/internal/model"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PaymentList struct {
	BaseCommand
}

func NewPaymentList(baseCmd *BaseCommand) *PaymentList {
	return &PaymentList{*baseCmd}
}

func (cmd *PaymentList) Name() string {
	return "list"
}

func (cmd *PaymentList) Description() string {
	return "Список всех платежей"
}

func (cmd *PaymentList) Handle(update tgbotapi.Update) error {
	userId := update.Message.From.ID
	payments, err := cmd.Service.Payment.Index(userId)

	if err != nil || len(payments) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не найдено платежей")

		_, err = cmd.TGBot.Send(msg)
		if err != nil {
			return err
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, formatPayments(payments))

	_, err = cmd.TGBot.Send(msg)

	return err
}

func formatPayments(payments []model.Payment) string {
	var message string

	for _, payment := range payments {
		message = fmt.Sprintf("%s \n"+payment.StringForTg(), message)
	}

	return message
}
