package command

import (
	"AutoPayment/internal/handler/telegram/action"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PaymentDelete struct {
	BaseCommand
}

func NewPaymentDelete(baseCmd *BaseCommand) *PaymentDelete {
	return &PaymentDelete{*baseCmd}
}

func (cmd *PaymentDelete) Name() string {
	return "delete"
}

func (cmd *PaymentDelete) Description() string {
	return "Удалить автоплатёж"
}

func (cmd *PaymentDelete) Handle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID
	message := "Введите полное название платежа, который хотите удалить"

	msg := tgbotapi.NewMessage(chatId, message)

	_, err := cmd.TGBot.Send(msg)

	return err
}

func (cmd *PaymentDelete) NextAction() *string {
	act := action.PaymentDelete{}
	name := act.Name()

	return &name
}
