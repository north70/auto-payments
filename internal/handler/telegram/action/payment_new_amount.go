package action

import (
	"AutoPayment/internal/handler/telegram/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type PaymentNewAmount struct {
	BaseAction
}

func NewPaymentNewAmount(baseAction *BaseAction) *PaymentNewAmount {
	return &PaymentNewAmount{BaseAction: *baseAction}
}

func (a *PaymentNewAmount) Name() string {
	return "payment_new_amount"
}

func (a *PaymentNewAmount) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	amountInFloat, err := strconv.ParseFloat(upd.Message.Text, 64)
	if err != nil {
		return errors.NewTgValidationError("Сумма платежа должна быть числом")
	}
	if amountInFloat == 0 {
		return errors.NewTgValidationError("Сумма платежа должна быть больше 0")
	}

	amountInInt := int(amountInFloat * 100)
	tempPayment.Amount = &amountInInt

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите кол-во платежей. Если платёж регулярный, то введите 0")
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentNewAmount) Next() string {
	next := PaymentNewCountPay{}

	return next.Name()
}
