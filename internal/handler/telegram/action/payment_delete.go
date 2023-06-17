package action

import (
	"AutoPayment/internal/handler/telegram/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PaymentDelete struct {
	BaseAction
}

func NewPaymentDelete(baseAction *BaseAction) *PaymentDelete {
	return &PaymentDelete{BaseAction: *baseAction}
}

func (a *PaymentDelete) Name() string {
	return "payment_delete"
}

func (a *PaymentDelete) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	name := upd.Message.Text

	err := a.Service.Payment.Delete(chatId, name)

	if err != nil {
		return errors.NewTgValidationError("Платёж с таким названием не найден. Попробуйте ещё раз")
	}

	msg := tgbotapi.NewMessage(chatId, "Платёж успешно удалён")
	_, err = a.TGBot.Send(msg)

	return err
}
