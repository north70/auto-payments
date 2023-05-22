package model

type Telegram struct {
	Id      int     `db:"id"`
	ChatId  int     `db:"chat_id"`
	Command string  `db:"command"`
	Action  *string `db:"action"`
}
