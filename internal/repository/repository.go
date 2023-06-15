package repository

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository/postgres"
	"AutoPayment/internal/repository/redis"
	"github.com/jmoiron/sqlx"
	redisClient "github.com/redis/go-redis/v9"
	"time"
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
	IndexByTime(limit, offset int, time time.Time) ([]model.Payment, error)
	Create(payment model.Payment) error
	IndexByChatId(chatId int64) ([]model.Payment, error)
	Show(id int) (model.Payment, error)
	Delete(id int) error
	Update(payment model.UpdatePayment) error
}

type Telegram interface {
	Upsert(chatId int64, command string, action *string) error
	UpdateAction(chatId int64, action *string) error
	Get(chatId int64) (model.Telegram, error)
}

type PaymentTemp interface {
	Flush(chatId int64) error
	Get(chatId int64) (model.PaymentTemp, error)
	SetOrUpdate(chatId int64, temp model.PaymentTemp) error
}
