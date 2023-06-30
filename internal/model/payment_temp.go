package model

type PaymentTemp struct {
	ID         *int
	ChatID     int64
	Name       *string
	PeriodDay  *int
	PaymentDay *int
	Amount     *int
	CountPay   *int
}

func (temp *PaymentTemp) ToMainStruct() Payment {
	mainS := Payment{
		ChatID:     temp.ChatID,
		Name:       *temp.Name,
		PeriodDay:  *temp.PeriodDay,
		PaymentDay: *temp.PaymentDay,
		Amount:     *temp.Amount,
		CountPay:   temp.CountPay,
	}

	return mainS
}

func (temp *PaymentTemp) ToUpdateStruct() UpdatePayment {
	update := UpdatePayment{
		ID:         *temp.ID,
		Name:       temp.Name,
		PeriodDay:  temp.PeriodDay,
		PaymentDay: temp.PaymentDay,
		Amount:     temp.Amount,
		CountPay:   temp.CountPay,
	}

	return update
}
