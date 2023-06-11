package action

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type PaymentNewName struct {
	BaseAction
}

func NewPaymentNewName(baseAction *BaseAction) *PaymentNewName {
	return &PaymentNewName{BaseAction: *baseAction}
}

func (a *PaymentNewName) Name() string {
	return "payment_new_name"
}

func (a *PaymentNewName) Handle(upd tgbotapi.Update) error {
	chatId := upd.Message.Chat.ID
	tempPayment, err := a.Service.PaymentTemp.Get(chatId)
	if err != nil {
		return err
	}
	name := upd.Message.Text
	tempPayment.Name = &name

	if err = a.Service.PaymentTemp.SetOrUpdate(chatId, tempPayment); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, "Введите периодичность платежа в днях")
	_, err = a.TGBot.Send(msg)

	return err
}

func (a *PaymentNewName) Next() string {
	next := PaymentNewPeriod{}

	return next.Name()
}
