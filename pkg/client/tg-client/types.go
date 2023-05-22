package tg_client

import "strings"

type Config struct {
	Token        string
	DialogEnable bool
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type UpdatesChannel <-chan Update

type Message struct {
	MessageId  int             `json:"message_id"`
	From       User            `json:"from"`
	SenderChat Chat            `json:"sender_chat"`
	Text       string          `json:"text"`
	Entities   []MessageEntity `json:"entities"`
}

type User struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Url    int    `json:"url"`
	User   User   `json:"user"`
}

func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	entity := m.Entities[0]
	return m.Text[1:entity.Length]
}

func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(m.Entities) == 0 {
		return false
	}

	entity := m.Entities[0]
	return entity.Offset == 0 && entity.IsCommand()
}

func (e MessageEntity) IsCommand() bool {
	return e.Type == "bot_command"
}

func (u Update) ChatId() int {
	return u.Message.From.Id
}
