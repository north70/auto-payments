package command

import (
	"AutoPayment/pkg/client/tg-client"
	"fmt"
)

type WhoamiCommand struct {
	Bot tg_client.BotApi
}

func NewWhoamiCommand(bot tg_client.BotApi) *WhoamiCommand {
	return &WhoamiCommand{bot}
}

func (cmd *WhoamiCommand) Name() string {
	return "whoami"
}

func (cmd *WhoamiCommand) IsDialog() bool {
	return false
}

func (cmd *WhoamiCommand) Description() string {
	return "Получить информацию о себе"
}
func (cmd *WhoamiCommand) Handle(update tg_client.Update) error {
	user := update.Message.From
	message := fmt.Sprintf("Имя - %s\nФамилия - %s\nUsername - %s\nТелеграм ID - %d",
		user.FirstName, user.LastName, user.Username, user.Id)

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: update.Message.From.Id,
		Text:   message,
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}
