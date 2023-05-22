package redis

import (
	"AutoPayment/internal/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var ctx = context.Background()

type PaymentTempRepository struct {
	db *redis.Client
}

func NewPaymentTempRepository(db *redis.Client) *PaymentTempRepository {
	return &PaymentTempRepository{db: db}
}

func (repo *PaymentTempRepository) Flush(chatId int) error {
	_, err := repo.db.Del(ctx, strconv.Itoa(chatId)).Result()

	return err
}

func (repo *PaymentTempRepository) Get(chatId int) (model.PaymentTemp, error) {
	data, err := repo.db.Get(ctx, strconv.Itoa(chatId)).Bytes()

	if err == redis.Nil {
		return model.PaymentTemp{}, errors.New("key does not exists")
	} else if err != nil {
		return model.PaymentTemp{}, err
	}

	tempP := model.PaymentTemp{}
	err = json.Unmarshal(data, &tempP)

	return tempP, err
}

func (repo *PaymentTempRepository) SetOrUpdate(chatId int, temp model.PaymentTemp) error {
	return nil
}
