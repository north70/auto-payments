package model

import (
	"fmt"
	"time"
)

type Payment struct {
	Id          int       `db:"id"`
	ChatId      int64     `db:"chat_id"`
	Name        string    `db:"name"`
	PeriodDay   int       `db:"period_day"`
	PaymentDay  int       `db:"payment_day"`
	Amount      int       `db:"amount"`
	CountPay    *int      `db:"count_pay"`
	NextPayDate time.Time `db:"next_pay_date"`
	CreatedAt   time.Time `db:"created_at"`
}

func (payment Payment) StringForTg() string {
	var payStr string

	amount := float32(payment.Amount) / 100
	payStr = fmt.Sprintf(payStr+
		"Название: %s\n"+
		"Периодичность платежа: %d дней\n"+
		"Следующтй платёж: %s\n"+
		"Сумма платежа: %.2f₽\n",
		payment.Name, payment.PeriodDay, payment.NextPayDate.Format("02.01.2006"), amount)

	if *payment.CountPay != 0 {
		payStr = fmt.Sprintf(payStr+"Кол-во платежей: %d\n", payment.CountPay)
	}

	return payStr
}

type UpdatePayment struct {
	Id          int
	Name        *string
	PeriodDay   *int
	PaymentDay  *int
	Amount      *int
	CountPay    *int
	NextPayDate *time.Time
}
