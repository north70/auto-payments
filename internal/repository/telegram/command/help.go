package command

import (
	"AutoPayment/pkg/client/tg-client"
	"fmt"
)

type HelpCommand struct {
	Bot tg_client.BotApi
}

func NewHelpCommand(bot tg_client.BotApi) *HelpCommand {
	return &HelpCommand{Bot: bot}
}

func (cmd *HelpCommand) Name() string {
	return "help"
}

func (cmd *HelpCommand) IsDialog() bool {
	return false
}

func (cmd *HelpCommand) Description() string {
	return "Получить все доступные команды"
}
func (cmd *HelpCommand) Handle(update tg_client.Update) error {
	commands := cmd.Bot.Commands.List()

	message := fmt.Sprintf("Список доступных комманды:\n/%s - %s", cmd.Name(), cmd.Description())

	for name, desc := range commands {
		message = fmt.Sprintf("%s\n/%s - %s", message, name, desc)
	}

	sendMsgQuery := tg_client.SendMessageQuery{
		ChatId: update.Message.From.Id,
		Text:   message,
	}

	err := cmd.Bot.SendMessage(sendMsgQuery)

	return err
}
