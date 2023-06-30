package command

import (
	"AutoPayment/internal/handler/telegram/action"
	"AutoPayment/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PaymentNew struct {
	BaseCommand
}

func NewPaymentNew(baseCmd *BaseCommand) *PaymentNew {
	return &PaymentNew{*baseCmd}
}

func (cmd *PaymentNew) Name() string {
	return "new"
}

func (cmd *PaymentNew) Description() string {
	return "Создать новый автоплатёж"
}

func (cmd *PaymentNew) Handle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID
	message := "Введите название нового платежа"

	paymentTemp := model.PaymentTemp{ChatID: chatId}
	err := cmd.Service.PaymentTemp.SetOrUpdate(chatId, paymentTemp)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, message)

	_, err = cmd.TGBot.Send(msg)

	return err
}

func (cmd *PaymentNew) NextAction() *string {
	act := action.PaymentNewName{}
	name := act.Name()

	return &name
}
