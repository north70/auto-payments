package action

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type PaymentNewPeriod struct {
	BaseAction
}

func NewPaymentNewPeriod(baseAction *BaseAction) *PaymentNewPeriod {
	return &PaymentNewPeriod{BaseAction: *baseAction}
}

func (a *PaymentNewPeriod) Name() string {
	return "payment_new_period"
}

func (a *PaymentNewPeriod) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	periodDay, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return errors.New("период платежа должен быть целым числом")
	}
	if periodDay < 1 || periodDay > 30 {
		return errors.New("период платежа может быть 1 до 30 дней")
	}

	fmt.Println(periodDay)
	tempPayment.PeriodDay = &periodDay

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите число месяца, когда проходит платеж")
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentNewPeriod) Next() string {
	next := PaymentNewDayPay{}

	return next.Name()
}
