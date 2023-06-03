package command

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/service"
	tg_client "AutoPayment/pkg/client/tg-client"
	"fmt"
	"strconv"
)

type PaymentNewCommand struct {
	Bot             tg_client.BotApi
	paymentTempServ service.PaymentTemp
	paymentServ     service.Payment
}

func NewPaymentNewCommand(bot tg_client.BotApi, paymentTempServ service.PaymentTemp, paymentServ service.Payment) *PaymentNewCommand {
	return &PaymentNewCommand{Bot: bot, paymentTempServ: paymentTempServ, paymentServ: paymentServ}
}

func (cmd *PaymentNewCommand) Name() string {
	return "new"
}

func (cmd *PaymentNewCommand) Description() string {
	return "Создать новый автоплатёж"
}

func (cmd *PaymentNewCommand) ActionList() map[string]func(upd tg_client.Update) error {

	return map[string]func(upd tg_client.Update) error{
		"name":     cmd.actionGetPeriodType,
		"period":   cmd.actionGetPeriod,
		"dayPay":   cmd.actionDayPay,
		"amount":   cmd.actionAmount,
		"countPay": cmd.actionCountPay,
		"create":   cmd.actionCreate,
	}
}

func (cmd *PaymentNewCommand) ActionMap() map[string]string {
	return map[string]string{
		"name":     "period",
		"period":   "dayPay",
		"dayPay":   "amount",
		"amount":   "countPay",
		"countPay": "create",
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

func (cmd *PaymentNewCommand) actionGetPeriodType(upd tg_client.Update) error {
	tempPayment := model.PaymentTemp{}
	name := upd.Message.Text
	chatId := upd.ChatId()

	tempPayment.Name = &name
	tempPayment.ChatId = &chatId

	if err := cmd.paymentTempServ.SetOrUpdate(upd.ChatId(), tempPayment); err != nil {
		return err
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите тип платежа. 1 - регулярный. 2 - временный",
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionGetPeriod(upd tg_client.Update) error {
	tempPayment, err := cmd.paymentTempServ.Get(upd.ChatId())
	if err != nil {
		return err
	}
	data, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return tg_client.NewMessageValidationError("Тип платежа должен быть целым числом")
	}
	if data != model.PeriodTypeRegular && data != model.PeriodTypeTemporary {
		return tg_client.NewMessageValidationError("Тип платежа может быть 1 - регулярный. 2 - временный")
	}

	periodType := model.PeriodType(data)

	tempPayment.PeriodType = &periodType

	if err = cmd.paymentTempServ.SetOrUpdate(upd.ChatId(), tempPayment); err != nil {
		return err
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите периодичность платежа в днях",
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionDayPay(upd tg_client.Update) error {
	tempPayment, err := cmd.paymentTempServ.Get(upd.ChatId())
	if err != nil {
		return err
	}
	periodDay, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return tg_client.NewMessageValidationError("Период платежа должен быть целым числом")
	}
	if periodDay < 1 || periodDay > 30 {
		return tg_client.NewMessageValidationError("Период платежа может быть 1 до 30 дней")
	}

	tempPayment.PeriodDay = &periodDay

	if err = cmd.paymentTempServ.SetOrUpdate(upd.ChatId(), tempPayment); err != nil {
		return err
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите число месяца, когда проходит платеж",
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionAmount(upd tg_client.Update) error {
	tempPayment, err := cmd.paymentTempServ.Get(upd.ChatId())
	if err != nil {
		return err
	}
	paymentDay, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return tg_client.NewMessageValidationError("Дата платежа должна быть целым числом")
	}
	if paymentDay < 1 || paymentDay > 30 {
		return tg_client.NewMessageValidationError("Дата платежа может быть числом от 1 до 30")
	}

	tempPayment.PaymentDay = &paymentDay

	if err = cmd.paymentTempServ.SetOrUpdate(upd.ChatId(), tempPayment); err != nil {
		return err
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите сумму платежа",
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionCountPay(upd tg_client.Update) error {
	tempPayment, err := cmd.paymentTempServ.Get(upd.ChatId())
	if err != nil {
		return err
	}
	amountInFloat, err := strconv.ParseFloat(upd.Message.Text, 64)
	if err != nil {
		return tg_client.NewMessageValidationError("Сумма платежа должна быть числом")
	}
	if amountInFloat == 0 {
		return tg_client.NewMessageValidationError("Сумма платежа должна быть больше 0")
	}

	amountInInt := int(amountInFloat * 100)

	tempPayment.Amount = &amountInInt

	if err = cmd.paymentTempServ.SetOrUpdate(upd.ChatId(), tempPayment); err != nil {
		return err
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Введите кол-во платежей. Если платёж регулярный, то введите 0",
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	return err
}

func (cmd *PaymentNewCommand) actionCreate(upd tg_client.Update) error {
	tempPayment, err := cmd.paymentTempServ.Get(upd.ChatId())
	if err != nil {
		return err
	}
	countPay, err := strconv.Atoi(upd.Message.Text)
	if err != nil {
		return tg_client.NewMessageValidationError("Кол-во платежей должно быть числом")
	}

	tempPayment.CountPay = &countPay
	tempPayment.IsFull = true

	if err = cmd.paymentTempServ.SetOrUpdate(upd.ChatId(), tempPayment); err != nil {
		return err
	}

	payment, err := tempPayment.ToMainStruct()
	if err != nil {
		return err
	}

	if err = cmd.paymentServ.Create(payment); err != nil {
		return err
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   "Авто-платёж успешно создан",
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	sendMsgQuery = tg_client.SendMessageQuery{
		ChatId: upd.Message.From.Id,
		Text:   payment.StringForTg(),
	}

	err = cmd.Bot.SendMessage(sendMsgQuery)

	return err
}
