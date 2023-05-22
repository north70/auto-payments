package tg_client

type Dialog interface {
	Command
	ActionList() map[string]func(upd Update) error
	ActionMap() map[string]string
	FirstAction() string
}

type ChatStatus struct {
	Command string
	Action  *string
}

type Store interface {
	Has(chatId int) (bool, error)
	New(chatId int, command string) error
	Current(chatId int) (ChatStatus, error)
	ClearAction(chatId int) error
	SetCommand(chatId int, command string) error
	SetAction(chatId int, action string) error
}
