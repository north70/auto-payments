package telegram

import (
	"AutoPayment/config"
	"AutoPayment/internal/handler/telegram/action"
	"AutoPayment/internal/handler/telegram/command"
	"AutoPayment/internal/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

type TgBot struct {
	Bot      *tgbotapi.BotAPI
	Config   config.Config
	Log      zerolog.Logger
	Service  *service.Service
	commands map[string]command.Command
	actions  map[string]action.Action
}

func NewTgBot(bot *tgbotapi.BotAPI, config config.Config, log zerolog.Logger, service *service.Service) *TgBot {
	return &TgBot{Bot: bot, Config: config, Log: log, Service: service}
}

func (t *TgBot) Start() {
	t.Bot.Debug = t.Config.BotDebug

	t.initActions()
	t.initCommands()

	t.Log.Info().Msg("telegram started...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = t.Config.BotTimout

	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			t.handleCommand(update)
		} else {
			t.handleMessage(update)
		}
	}
}

func (t *TgBot) initCommands() {
	baseCmd := command.NewBaseCommand(t.Bot, t.Service, t.actions)

	commands := t.appendCommands([]command.Command{
		command.NewWhoami(baseCmd),
		command.NewPaymentList(baseCmd),
		command.NewPaymentNew(baseCmd),
	})

	t.commands = commands
}

func (t *TgBot) appendCommands(commands []command.Command) map[string]command.Command {
	cmdList := make(map[string]command.Command, len(commands))

	for _, cmd := range commands {
		cmdList[cmd.Name()] = cmd
	}

	return cmdList
}

func (t *TgBot) initActions() {
	baseAction := action.NewBaseAction(t.Bot, t.Service)

	actions := t.appendActions([]action.Action{
		action.NewPaymentNewName(baseAction),
		action.NewPaymentNewPeriodType(baseAction),
		action.NewPaymentNewAmount(baseAction),
		action.NewPaymentNewCountPay(baseAction),
		action.NewPaymentNewDayPay(baseAction),
		action.NewPaymentNewPeriod(baseAction),
	})

	t.actions = actions
}

func (t *TgBot) appendActions(actions []action.Action) map[string]action.Action {
	actionList := make(map[string]action.Action, len(actions))

	for _, action := range actions {
		actionList[action.Name()] = action
	}

	return actionList
}

func (t *TgBot) handleCommand(upd tgbotapi.Update) {
	cmdName := upd.Message.Command()

	cmd, ok := t.commands[cmdName]
	if !ok {
		t.Log.Debug().Msg(fmt.Sprintf("command '%s' not found", cmdName))
		return
	}

	chatId := upd.Message.Chat.ID
	err := cmd.Handle(upd)
	if err != nil {
		t.Log.Err(err).Msg(fmt.Sprintf("error handle command %s for chat %d", cmdName, chatId))
		return
	}

	err = t.Service.Telegram.Upsert(chatId, cmdName, cmd.NextAction())
	if err != nil {
		t.Log.Err(err).Msg(fmt.Sprintf("error update chat %d", chatId))
	}
}

func (t *TgBot) handleMessage(upd tgbotapi.Update) {
	chatId := upd.Message.Chat.ID

	chat, err := t.Service.Telegram.Get(chatId)
	if err != nil {
		t.Log.Debug().Msg(fmt.Sprintf("chat %d not found", chatId))
		return
	}

	if chat.Action == nil {
		return
	}

	act, ok := t.actions[*chat.Action]
	if !ok {
		t.Log.Debug().Msg(fmt.Sprintf("action '%s' not found", *chat.Action))
		return
	}

	if err = act.Handle(upd); err != nil {
		t.Log.Err(err).Msg(fmt.Sprintf("error handle action '%s' for chat %d", act.Name(), chatId))
		return
	}

	var nextAct *string
	if name := act.Next(); name != "" {
		nextAct = &name
	}
	if err = t.Service.Telegram.UpdateAction(chatId, nextAct); err != nil {
		t.Log.Err(err).Msg(fmt.Sprintf("error update chat %d", chatId))
		return
	}
}
