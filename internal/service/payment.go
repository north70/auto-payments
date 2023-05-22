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

func (s *PaymentService) Index(userId int) ([]model.Payment, error) {
	return s.repo.Index(userId)
}

func (s *PaymentService) Show(userId, id int) (model.Payment, error) {
	return s.repo.Show(userId, id)
}

func (s *PaymentService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
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

func (s *PaymentTempService) Flush(chatId int) error {
	return s.repo.Flush(chatId)
}

func (s *PaymentTempService) Get(chatId int) (model.PaymentTemp, error) {
	return s.repo.Get(chatId)
}

func (s *PaymentTempService) SetOrUpdate(chatId int, temp model.PaymentTemp) error {
	return s.repo.SetOrUpdate(chatId, temp)
}
