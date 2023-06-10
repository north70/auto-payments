package command

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Whoami struct {
	BaseCommand
}

func NewWhoami(baseCmd *BaseCommand) *Whoami {
	return &Whoami{*baseCmd}
}

func (cmd *Whoami) Name() string {
	return "whoami"
}

func (cmd *Whoami) Description() string {
	return "Получить информацию о себе"
}
func (cmd *Whoami) Handle(update tgbotapi.Update) error {
	user := update.Message.From
	message := fmt.Sprintf("Имя - %s\nФамилия - %s\nUsername - %s\nТелеграм ID - %d",
		user.FirstName, user.LastName, user.UserName, user.ID)

	msg := tgbotapi.NewMessage(user.ID, message)

	_, err := cmd.TGBot.Send(msg)

	return err
}
