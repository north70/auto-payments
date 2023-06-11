package action

import (
	"AutoPayment/internal/handler/telegram/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type PaymentNewCountPay struct {
	BaseAction
}

func NewPaymentNewCountPay(baseAction *BaseAction) *PaymentNewCountPay {
	return &PaymentNewCountPay{BaseAction: *baseAction}
}

func (a *PaymentNewCountPay) Name() string {
	return "payment_new_count_pay"
}

func (a *PaymentNewCountPay) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	countPay, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return errors.NewTgValidationError("Кол-во платежей должно быть числом")
	}

	tempPayment.CountPay = &countPay
	tempPayment.IsFull = true

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}
	payment, err := tempPayment.ToMainStruct()
	if err != nil {
		return err
	}

	if err = a.Service.Payment.Create(payment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Авто-платёж успешно создан")
	_, err = a.TGBot.Send(msg)
	if err != nil {
		return err
	}

	msg = tgbotapi.NewMessage(chatId, payment.StringForTg())
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentNewCountPay) Next() string {
	return ""
}
