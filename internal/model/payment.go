package model

import (
	"fmt"
	"time"
)

type Payment struct {
	Id         int        `db:"id"`
	ChatId     int64      `db:"chat_id"`
	Name       string     `db:"name"`
	PeriodType PeriodType `db:"period_type"`
	PeriodDay  int        `db:"period_day"`
	PaymentDay int        `db:"payment_day"`
	Amount     int        `db:"amount"`
	CountPay   *int       `db:"count_pay"`
	CreatedAt  time.Time  `db:"created_at"`
}

func (payment Payment) StringForTg() string {
	var payStr string
	if payment.Id != 0 {
		payStr = fmt.Sprintf("ID: %d\n", payment.Id)
	}

	amount := float32(payment.Amount) / 100
	payStr = fmt.Sprintf(payStr+
		"Название: %s\n"+
		"Пероидичность платежа: %d\n"+
		"Число платежей: %d\n"+
		"Сумма платежа: %.2f\n",
		payment.Name, payment.PeriodDay, payment.PaymentDay, amount)

	if payment.PeriodType == PeriodTypeTemporary {
		payStr = fmt.Sprintf(payStr+"Кол-во платежей: %d\n", payment.CountPay)
	}

	return payStr
}

type UpdatePayment struct {
	Id         int
	ChatId     int
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
