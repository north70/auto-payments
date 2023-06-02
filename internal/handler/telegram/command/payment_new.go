package command

import (
	tg_client "AutoPayment/pkg/client/tg-client"
	"fmt"
)

type PaymentNewCommand struct {
	Bot tg_client.BotApi
}

func NewPaymentNewCommand(bot tg_client.BotApi) *PaymentNewCommand {
	return &PaymentNewCommand{Bot: bot}
}

func (cmd *PaymentNewCommand) Name() string {
	return "new"
}

func (cmd *PaymentNewCommand) Description() string {
	return "Создать новый автоплатёж"
}

func (cmd *PaymentNewCommand) ActionList() map[string]func(upd tg_client.Update) error {

	return map[string]func(upd tg_client.Update) error{
		"name":     cmd.actionGetName,
		"period":   cmd.actionGetPeriod,
		"dayPay":   cmd.actionDayPay,
		"amount":   cmd.actionAmount,
		"countPay": cmd.actionCountPay,
	}
}

func (cmd *PaymentNewCommand) ActionMap() map[string]string {
	return map[string]string{
		"name":   "period",
		"period": "dayPay",
		"dayPay": "amount",
		"amount": "countPay",
	}
}

func (cmd *PaymentNewCommand) FirstAction() string {
	return "name"
}

func (cmd *PaymentNewCommand) Handle(update tg_client.Update) error {

	message := fmt.Sprintf("Введите название нового платежа")

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: update.Message.From.Id,
		Text:   message,
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionGetName(upd tg_client.Update) error {

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите тип платежа",
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionGetPeriod(upd tg_client.Update) error {
	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите период платежа",
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionDayPay(upd tg_client.Update) error {
	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите дату платежа",
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionAmount(upd tg_client.Update) error {
	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите сумму платежа",
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionCountPay(upd tg_client.Update) error {
	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите кол-во платежей",
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}
