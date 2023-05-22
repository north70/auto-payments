package model

import (
	"errors"
	"time"
)

type Payment struct {
	Id         int        `db:"id"`
	ChatId     int        `db:"chat_id"`
	Name       string     `db:"name"`
	PeriodType PeriodType `db:"period_type"`
	PeriodDay  int        `db:"period_day"`
	PaymentDay int        `db:"payment_day"`
	Amount     int        `db:"amount"`
	CountPay   *int       `db:"count_pay"`
	CreatedAt  time.Time  `db:"created_at"`
}

type PaymentTemp struct {
	IsFull     bool
	ChatId     *int
	Name       *string
	PeriodType *PeriodType
	PeriodDay  *int
	PaymentDay *int
	Amount     *int
	CountPay   *int
}

func (temp PaymentTemp) ToMainStruct() (Payment, error) {
	if !temp.IsFull {
		return Payment{}, errors.New("struct is not full")
	}

	mainS := Payment{
		ChatId:     *temp.ChatId,
		Name:       *temp.Name,
		PeriodType: *temp.PeriodType,
		PeriodDay:  *temp.PeriodDay,
		PaymentDay: *temp.PaymentDay,
		Amount:     *temp.Amount,
		CountPay:   temp.CountPay,
	}

	return mainS, nil
}

type UpdatePayment struct {
	Id         int
	UserId     int
	Name       *string
	PeriodType *PeriodType
	PeriodDay  *int
	PaymentDay *int
	Amount     *int
	CountPay   *int
	CreatedAt  time.Time
}

type PeriodType int

const (
	PeriodTypeRegular = iota + 1
	PeriodTypeTemporary
)
