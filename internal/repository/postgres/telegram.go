package postgres

import (
	"AutoPayment/internal/model"
	"github.com/jmoiron/sqlx"
)

type TelegramRepository struct {
	db *sqlx.DB
}

func NewTelegramRepository(db *sqlx.DB) *TelegramRepository {
	return &TelegramRepository{db: db}
}

func (repo *TelegramRepository) Get(chatId int64) (model.Telegram, error) {
	data := model.Telegram{}

	query := "SELECt * FROM telegram WHERE chat_id = $1"
	err := repo.db.Get(&data, query, chatId)
	if err != nil {
		return model.Telegram{}, err
	}

	return data, nil
}

func (repo *TelegramRepository) UpdateAction(chatId int64, action *string) error {
	query := "UPDATE telegram SET action = $1 WHERE chat_id = $2"

	_, err := repo.db.Exec(query, action, chatId)

	return err
}

func (repo *TelegramRepository) Upsert(chatId int64, command string, action *string) error {
	query := `INSERT INTO telegram (chat_id, command, action) VALUES ($1, $2, $3)
			  ON CONFLICT (chat_id) DO UPDATE SET command = $2, action = $3`

	_, err := repo.db.Exec(query, chatId, command, action)

	return err
}
