package service

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository"
)

type Service struct {
	Payment
	PaymentTemp
	Telegram
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Payment:     NewPaymentService(repo.Payment),
		PaymentTemp: NewPaymentTempService(repo.PaymentTemp),
		Telegram:    NewTelegramService(repo.Telegram),
	}
}

type Messenger interface {
	SendMessage(recipientId string, message string) error
	GetLastMessage(senderId string, message string) (string, error)
}

type Payment interface {
	Create(payment model.Payment) error
	Index(chatId int64) ([]model.Payment, error)
	Show(chatId int64, id int) (model.Payment, error)
	Delete(chatId int64, id int) error
	Update(payment model.UpdatePayment) error
}

type PaymentTemp interface {
	Flush(chatId int64) error
	Get(chatId int64) (model.PaymentTemp, error)
	SetOrUpdate(chatId int64, temp model.PaymentTemp) error
}

type Telegram interface {
	Get(chatId int64) (model.Telegram, error)
	Upsert(chatId int64, command string, action *string) error
	UpdateAction(chatId int64, action *string) error
}
