package tg_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const baseUri = "https://api.telegram.org/bot"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type BotApi struct {
	Config          Config
	Client          HttpClient
	Commands        Commands
	Store           Store
	shutdownChannel chan interface{}
}

type QueryParams map[string]string

type ApiResponse struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

func NewBotApi(cfg Config) *BotApi {
	return &BotApi{Config: cfg, Client: &http.Client{}}
}

func (bot *BotApi) MakeRequest(method string, params QueryParams) (*ApiResponse, error) {
	values := buildParams(params)

	fullUrl := fmt.Sprintf("%s%s/%s?%s", baseUri, bot.Config.Token, method, values.Encode())

	req, err := http.NewRequest(http.MethodGet, fullUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return &ApiResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse

	err = bot.decodeApiResponse(resp.Body, &apiResponse)
	if err != nil {
		return &ApiResponse{}, err
	}

	return &apiResponse, nil
}

func (bot *BotApi) GetUpdates(query UpdateQuery) ([]Update, error) {
	params := query.params()

	resp, err := bot.MakeRequest("getUpdates", params)
	if err != nil {
		return []Update{}, err
	}

	var updates []Update
	err = json.Unmarshal(resp.Result, &updates)

	return updates, err
}

func (bot *BotApi) GetUpdatesChan(query UpdateQuery) UpdatesChannel {
	ch := make(chan Update)

	go func() {
		for {
			select {
			case <-bot.shutdownChannel:
				close(ch)
				return
			default:
			}

			updates, err := bot.GetUpdates(query)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateId >= query.Offset {
					query.Offset = update.UpdateId + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func (bot *BotApi) SendMessage(query SendMessageQuery) error {
	params := query.params()

	_, err := bot.MakeRequest("sendMessage", params)

	return err
}

func (bot *BotApi) HandleMessage(update Update) error {
	if commandName := update.Message.Command(); commandName != "" {
		err := bot.handleCommand(commandName, update)
		return err
	}

	if !bot.Config.DialogEnable {
		return nil
	}

	return bot.handleAction(update)
}

func (bot *BotApi) handleCommand(commandName string, upd Update) error {
	if !bot.Commands.has(commandName) {
		return errors.New(fmt.Sprintf("not found command %s", commandName))
	}

	cmd := bot.Commands.Commands[commandName]

	err := cmd.Handle(upd)
	if err != nil {
		return err
	}
	chatId := upd.ChatId()

	hasInStore, err := bot.Store.Has(chatId)

	if err != nil {
		return err
	}

	if hasInStore {
		err = bot.Store.SetCommand(chatId, cmd.Name())
	} else {
		err = bot.Store.New(chatId, cmd.Name())
	}

	if err != nil {
		return err
	}

	if !bot.Config.DialogEnable {
		return nil
	}

	dialog, ok := cmd.(Dialog)
	if !ok {
		return nil
	}

	return bot.Store.SetAction(chatId, dialog.FirstAction())
}

func (bot *BotApi) handleAction(upd Update) error {
	chatId := upd.ChatId()

	current, err := bot.Store.Current(chatId)
	if err != nil {
		return err
	}

	dialog, ok := bot.Commands.Commands[current.Command].(Dialog)
	if !ok {
		return nil
	}

	if current.Action == nil {
		message := fmt.Sprintf("Неизвестная операция")

		sendMsgQuery := SendMessageQuery{
			ChatId: chatId,
			Text:   message,
		}

		if err := bot.SendMessage(sendMsgQuery); err != nil {
			return err
		}
		return nil
	}

	actionName := *current.Action

	if err := dialog.ActionList()[actionName](upd); err != nil {
		return err
	}

	nextAction, ok := dialog.ActionMap()[actionName]
	if !ok {
		err = bot.Store.ClearAction(chatId)
	} else {
		err = bot.Store.SetAction(chatId, nextAction)
	}

	return err
}

func (bot *BotApi) decodeApiResponse(body io.Reader, apiResp *ApiResponse) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, apiResp)
	if err != nil {
		return err
	}

	return nil
}

func buildParams(in QueryParams) url.Values {
	if in == nil {
		return url.Values{}
	}

	out := url.Values{}

	for key, value := range in {
		out.Set(key, value)
	}

	return out
}
