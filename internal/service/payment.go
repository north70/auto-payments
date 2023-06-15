package service

import (
	"AutoPayment/internal/model"
	"AutoPayment/internal/repository"
	"time"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(payment model.Payment) error {
	today := time.Now()
	var nextPayDay time.Time
	if payment.PaymentDay > today.Day() {
		nextPayDay = today.AddDate(0, 0, payment.PaymentDay-today.Day())
	} else {
		nextPayDay = today.AddDate(0, 1, payment.PaymentDay-today.Day())
	}
	payment.NextPayDate = nextPayDay

	return s.repo.Create(payment)
}

func (s *PaymentService) IndexByChatId(chatId int64) ([]model.Payment, error) {
	return s.repo.IndexByChatId(chatId)
}

func (s *PaymentService) Show(id int) (model.Payment, error) {
	return s.repo.Show(id)
}

func (s *PaymentService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PaymentService) Update(payment model.UpdatePayment) error {
	return s.repo.Update(payment)
}

func (s *PaymentService) IndexByTime(limit, offset int, time time.Time) ([]model.Payment, error) {
	return s.repo.IndexByTime(limit, offset, time)
}

func (s *PaymentService) UpdateNextPayDay(id int) error {
	payment, err := s.Show(id)
	if err != nil {
		return err
	}

	now := time.Now()

	if payment.NextPayDate.After(now) {
		return nil
	}

	nextPayDate := payment.NextPayDate
	for {
		nextPayDate = nextPayDate.AddDate(0, 0, payment.PeriodDay)
		if nextPayDate.After(now) {
			break
		}
	}

	upd := model.UpdatePayment{
		Id:          id,
		NextPayDate: &nextPayDate,
	}
	err = s.Update(upd)

	return err
}
