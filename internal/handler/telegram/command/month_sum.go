package command

import (
	"AutoPayment/internal/handler/telegram/errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sum struct {
	BaseCommand
}

func NewSum(baseCmd *BaseCommand) *Sum {
	return &Sum{*baseCmd}
}

func (cmd *Sum) Name() string {
	return "sum"
}

func (cmd *Sum) Description() string {
	return "Получить сумму платежей за месяц"
}
func (cmd *Sum) Handle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID

	sumInCoin, err := cmd.Service.Payment.SumForMonth(chatId)
	if err != nil {
		return errors.NewTgValidationError("У вас нет активных платежей")
	}
	sumWithCoin := float64(sumInCoin) / 100

	message := fmt.Sprintf("Сумма платежей за месяц %.2f₽", sumWithCoin)

	msg := tgbotapi.NewMessage(chatId, message)

	_, err = cmd.TGBot.Send(msg)

	return err
}
