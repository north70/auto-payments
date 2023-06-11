package action

import (
	"AutoPayment/internal/handler/telegram/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type PaymentNewDayPay struct {
	BaseAction
}

func NewPaymentNewDayPay(baseAction *BaseAction) *PaymentNewDayPay {
	return &PaymentNewDayPay{BaseAction: *baseAction}
}

func (a *PaymentNewDayPay) Name() string {
	return "payment_new_day_pay"
}

func (a *PaymentNewDayPay) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	paymentDay, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return errors.NewTgValidationError("Дата платежа должна быть целым числом")
	}
	if paymentDay < 1 || paymentDay > 30 {
		return errors.NewTgValidationError("Дата платежа может быть числом от 1 до 30")
	}

	tempPayment.PaymentDay = &paymentDay

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите сумму платежа")
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentNewDayPay) Next() string {
	next := PaymentNewAmount{}

	return next.Name()
}
