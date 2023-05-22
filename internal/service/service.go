package service

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository"
	tg_client "AutoPayment/pkg/client/tg-client"
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
	Index(userId int) ([]model.Payment, error)
	Show(userId, id int) (model.Payment, error)
	Delete(userId, id int) error
	Update(payment model.UpdatePayment) error
}

type PaymentTemp interface {
	Flush(chatId int) error
	Get(chatId int) (model.PaymentTemp, error)
	SetOrUpdate(chatId int, temp model.PaymentTemp) error
}

type Telegram interface {
	Has(chatId int) (bool, error)
	New(chatId int, command string) error
	Current(chatId int) (tg_client.ChatStatus, error)
	ClearAction(chatId int) error
	SetCommand(chatId int, command string) error
	SetAction(chatId int, action string) error
}
