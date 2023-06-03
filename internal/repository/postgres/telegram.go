package postgres

import (
	"AutoPayment/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TelegramRepository struct {
	db *sqlx.DB
}

func NewTelegramRepository(db *sqlx.DB) *TelegramRepository {
	return &TelegramRepository{db: db}
}

func (repo *TelegramRepository) Create(chatId int, command string) error {
	query := "INSERT INTO telegram (chat_id, command) VALUES ($1, $2)"

	_, err := repo.db.Exec(query, chatId, command)

	return err
}

func (repo *TelegramRepository) Update(chatId int, command, action *string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	numParam := 1

	if command != nil {
		setValues = append(setValues, fmt.Sprintf("command=$%d", numParam))
		args = append(args, *command)
		numParam++
	}

	setValues = append(setValues, fmt.Sprintf("action=$%d", numParam))
	args = append(args, action, chatId)
	numParam++

	values := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE telegram SET %s WHERE chat_id = $%d", values, numParam)
	_, err := repo.db.Exec(query, args...)

	return err
}

func (repo *TelegramRepository) Get(chatId int) (model.Telegram, error) {
	data := model.Telegram{}

	query := "SELECt * FROM telegram WHERE chat_id = $1"
	err := repo.db.Get(&data, query, chatId)
	if err != nil {
		return model.Telegram{}, err
	}

	return data, nil
}

func (repo *TelegramRepository) Has(chatId int) (bool, error) {
	var id int

	query := "SELECt id FROM telegram WHERE chat_id = $1"
	result, err := repo.db.Query(query, chatId)
	if err != nil {
		return false, err
	}

	for result.Next() {
		err := result.Scan(&id)
		if err != nil {
			return false, err
		}
	}

	return id != 0, nil
}

func (repo *TelegramRepository) ClearAction(chatId int) error {
	query := "UPDATE telegram SET action = $1 WHERE chat_id = $2"

	_, err := repo.db.Exec(query, nil, chatId)

	return err
}
