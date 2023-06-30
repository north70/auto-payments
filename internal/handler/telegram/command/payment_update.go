package command

import (
	"AutoPayment/internal/handler/telegram/action"
	"AutoPayment/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PaymentUpdate struct {
	BaseCommand
}

func NewPaymentUpdate(baseCmd *BaseCommand) *PaymentUpdate {
	return &PaymentUpdate{*baseCmd}
}

func (cmd *PaymentUpdate) Name() string {
	return "update"
}

func (cmd *PaymentUpdate) Description() string {
	return "Обновить автоплатёж"
}

func (cmd *PaymentUpdate) Handle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID

	paymentTemp := model.PaymentTemp{ChatID: chatId}
	err := cmd.Service.PaymentTemp.SetOrUpdate(chatId, paymentTemp)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите полное название платежа, который хотите изменить")

	_, err = cmd.TGBot.Send(msg)

	return err
}

func (cmd *PaymentUpdate) NextAction() *string {
	act := action.PaymentUpdate{}
	name := act.Name()

	return &name
}
