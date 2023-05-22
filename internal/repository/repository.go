package repository

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository/postgres"
	"AutoPayment/internal/repository/redis"
	"github.com/jmoiron/sqlx"
	redisClient "github.com/redis/go-redis/v9"
)

type Repository struct {
	Payment
	PaymentTemp
	Telegram
}

func NewRepository(pgDB *sqlx.DB, cacheDB *redisClient.Client) *Repository {
	return &Repository{
		Payment:     postgres.NewPaymentRepository(pgDB),
		PaymentTemp: redis.NewPaymentTempRepository(cacheDB),
		Telegram:    postgres.NewTelegramRepository(pgDB),
	}
}

type Payment interface {
	Create(payment model.Payment) error
	Index(userId int) ([]model.Payment, error)
	Show(userId, id int) (model.Payment, error)
	Delete(userId, id int) error
	Update(payment model.UpdatePayment) error
}

type Telegram interface {
	Create(chatId int, command string) error
	Has(chatId int) (bool, error)
	Update(chatId int, command, action *string) error
	Get(chatId int) (model.Telegram, error)
	ClearAction(chatId int) error
}

type PaymentTemp interface {
	Flush(chatId int) error
	Get(chatId int) (model.PaymentTemp, error)
	SetOrUpdate(chatId int, temp model.PaymentTemp) error
}
