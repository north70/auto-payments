package model

type Telegram struct {
	ID      int     `db:"id"`
	ChatID  int     `db:"chat_id"`
	Command string  `db:"command"`
	Action  *string `db:"action"`
}
