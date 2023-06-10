package redis

import (
	"AutoPayment/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var ctx = context.Background()

const cacheTTL = 120

type PaymentTempRepository struct {
	db *redis.Client
}

func NewPaymentTempRepository(db *redis.Client) *PaymentTempRepository {
	return &PaymentTempRepository{db: db}
}

func (repo *PaymentTempRepository) Flush(chatId int64) error {
	_, err := repo.db.Del(ctx, strconv.FormatInt(chatId, 10)).Result()

	return err
}

func (repo *PaymentTempRepository) Get(chatId int64) (model.PaymentTemp, error) {
	data, err := repo.db.Get(ctx, strconv.FormatInt(chatId, 10)).Bytes()

	if err == redis.Nil {
		return model.PaymentTemp{}, errors.New("key does not exists")
	} else if err != nil {
		return model.PaymentTemp{}, err
	}

	tempP := model.PaymentTemp{}
	err = json.Unmarshal(data, &tempP)

	return tempP, err
}

func (repo *PaymentTempRepository) SetOrUpdate(chatId int64, temp model.PaymentTemp) error {
	jsonData, err := json.Marshal(temp)
	if err != nil {
		return errors.New("error encode temp model into json")
	}

	_, err = repo.db.Set(ctx, strconv.FormatInt(chatId, 10), jsonData, cacheTTL*time.Second).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("error set cache for chat_id = %d", chatId))
	}

	return nil
}
