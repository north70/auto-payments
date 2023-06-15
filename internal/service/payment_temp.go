package service

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository"
)

type PaymentTempService struct {
	repo repository.PaymentTemp
}

func NewPaymentTempService(repo repository.PaymentTemp) *PaymentTempService {
	return &PaymentTempService{repo: repo}
}

func (s *PaymentTempService) Flush(chatId int64) error {
	return s.repo.Flush(chatId)
}

func (s *PaymentTempService) Get(chatId int64) (model.PaymentTemp, error) {
	return s.repo.Get(chatId)
}

func (s *PaymentTempService) SetOrUpdate(chatId int64, temp model.PaymentTemp) error {
	return s.repo.SetOrUpdate(chatId, temp)
}
