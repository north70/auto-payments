package service

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(payment model.Payment) error {
	return s.repo.Create(payment)
}

func (s *PaymentService) Index(chatId int64) ([]model.Payment, error) {
	return s.repo.Index(chatId)
}

func (s *PaymentService) Show(chatId int64, id int) (model.Payment, error) {
	return s.repo.Show(chatId, id)
}

func (s *PaymentService) Delete(chatId int64, id int) error {
	return s.repo.Delete(chatId, id)
}

func (s *PaymentService) Update(payment model.UpdatePayment) error {
	return s.repo.Update(payment)
}

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
