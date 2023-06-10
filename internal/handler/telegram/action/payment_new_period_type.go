package action

import (
	"AutoPayment/internal/model"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type PaymentNewPeriodType struct {
	BaseAction
}

func NewPaymentNewPeriodType(baseAction *BaseAction) *PaymentNewPeriodType {
	return &PaymentNewPeriodType{BaseAction: *baseAction}
}

func (a *PaymentNewPeriodType) Name() string {
	return "payment_new_period_type"
}

func (a *PaymentNewPeriodType) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	data, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return errors.New("тип платежа должен быть целым числом")
	}
	if data != model.PeriodTypeRegular && data != model.PeriodTypeTemporary {
		return errors.New("тип платежа может быть 1 - регулярный. 2 - временный")
	}

	periodType := model.PeriodType(data)

	tempPayment.PeriodType = &periodType

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите периодичность платежа в днях")
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentNewPeriodType) Next() string {
	next := PaymentNewPeriod{}

	return next.Name()
}
