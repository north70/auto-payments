package action

import (
	"AutoPayment/internal/handler/telegram/errors"
	"AutoPayment/internal/model"
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

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	var payment model.Payment
	if tempPayment.ID == nil {
		paymentMain := tempPayment.ToMainStruct()
		payment, err = a.Service.Payment.Create(paymentMain)
	} else {
		paymentUpd := tempPayment.ToUpdateStruct()
		nextPayDate := model.CalcFirstNextPayDate(*paymentUpd.PaymentDay)
		paymentUpd.NextPayDate = &nextPayDate
		payment, err = a.Service.Payment.Update(paymentUpd)
	}
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Авто-платёж успешно создан/обновлён")
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
