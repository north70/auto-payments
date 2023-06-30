package model

import (
	"fmt"
	"time"
)

type Payment struct {
	ID          int       `db:"id"`
	ChatID      int64     `db:"chat_id"`
	Name        string    `db:"name"`
	PeriodDay   int       `db:"period_day"`
	PaymentDay  int       `db:"payment_day"`
	Amount      int       `db:"amount"`
	CountPay    *int      `db:"count_pay"`
	NextPayDate time.Time `db:"next_pay_date"`
	CreatedAt   time.Time `db:"created_at"`
}

func (p *Payment) StringForTg() string {
	var payStr string

	amount := float32(p.Amount) / 100
	payStr = fmt.Sprintf(payStr+
		"Название: %s\n"+
		"Периодичность платежа: %d дней\n"+
		"Следующтй платёж: %s\n"+
		"Сумма платежа: %.2f₽\n",
		p.Name, p.PeriodDay, p.NextPayDate.Format("02.01.2006"), amount)

	if *p.CountPay != 0 {
		payStr = fmt.Sprintf(payStr+"Кол-во платежей: %d\n", p.CountPay)
	}

	return payStr
}

func CalcFirstNextPayDate(paymentDay int) time.Time {
	today := time.Now()
	var nextPayDay time.Time
	if paymentDay > today.Day() {
		nextPayDay = today.AddDate(0, 0, paymentDay-today.Day())
	} else {
		nextPayDay = today.AddDate(0, 1, paymentDay-today.Day())
	}

	return nextPayDay
}

type UpdatePayment struct {
	ID          int
	Name        *string
	PeriodDay   *int
	PaymentDay  *int
	Amount      *int
	CountPay    *int
	NextPayDate *time.Time
}
