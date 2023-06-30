package action

import (
	"AutoPayment/internal/handler/telegram/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PaymentUpdate struct {
	BaseAction
}

func NewPaymentUpdate(baseAction *BaseAction) *PaymentUpdate {
	return &PaymentUpdate{BaseAction: *baseAction}
}

func (a *PaymentUpdate) Name() string {
	return "payment_update_name"
}

func (a *PaymentUpdate) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	name := upd.Message.Text
	tempPayment.Name = &name

	exists, err := a.Service.Payment.ExistsByName(chatId, name)
	if err != nil {
		return err
	}

	if !exists {
		return errors.NewTgValidationError("Платёж с таким название не найден")
	}

	payment, err := a.Service.Payment.GetByName(chatId, name)
	if err != nil {
		return err
	}

	tempPayment.ID = &payment.ID
	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите новое название платежа")
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentUpdate) Next() string {
	next := PaymentNewName{}

	return next.Name()
}
