package model

import "errors"

type PaymentTemp struct {
	IsFull     bool
	ChatId     int64
	Name       *string
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
		ChatId:     temp.ChatId,
		Name:       *temp.Name,
		PeriodDay:  *temp.PeriodDay,
		PaymentDay: *temp.PaymentDay,
		Amount:     *temp.Amount,
		CountPay:   temp.CountPay,
	}

	return mainS, nil
}
