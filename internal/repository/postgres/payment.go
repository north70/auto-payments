package postgres

import (
	"AutoPayment/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payment model.Payment) error {
	query := `INSERT INTO auto_payments (chat_id, name, period_day, payment_day, amount, count_pay, next_pay_date, created_at) 
			  VALUES (:chat_id, :name, :period_day, :payment_day, :amount, :count_pay, :next_pay_date, :created_at)`
	payment.CreatedAt = time.Now()

	_, err := r.db.NamedExec(query, payment)

	return err
}

func (r *PaymentRepository) Index(chatId int64) ([]model.Payment, error) {
	var models []model.Payment
	query := "SELECT * FROM auto_payments WHERE chat_id = $1"

	err := r.db.Select(&models, query, chatId)

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (r *PaymentRepository) Show(chatId int64, id int) (model.Payment, error) {
	payment := model.Payment{}
	query := "SELECT * FROM auto_payments WHERE id = $1 and chat_id = $2"

	err := r.db.Get(&payment, query, id, chatId)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *PaymentRepository) Delete(chatId int64, id int) error {
	query := fmt.Sprintf("DELETE FROM auto_payments WHERE chat_id = $1 and id = $2")
	_, err := r.db.Exec(query, chatId, id)

	return err
}

func (r *PaymentRepository) Update(payment model.UpdatePayment) error {
	args := make([]interface{}, 0)
	setValues := make([]string, 0)
	numParam := 1

	if payment.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", numParam))
		args = append(args, payment.Name)
		numParam++
	}

	if payment.PeriodDay != nil {
		setValues = append(setValues, fmt.Sprintf("period_day = $%d", numParam))
		args = append(args, payment.PeriodDay)
		numParam++
	}

	if payment.PaymentDay != nil {
		setValues = append(setValues, fmt.Sprintf("payment_day = $%d", numParam))
		args = append(args, payment.PaymentDay)
		numParam++
	}

	if payment.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount = $%d", numParam))
		args = append(args, payment.Name)
		numParam++
	}

	if payment.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("next_pay_date = $%d", numParam))
		args = append(args, payment.NextPayDate)
		numParam++
	}

	setValues = append(setValues, fmt.Sprintf("count_pay = $%d", numParam))
	numParam++

	args = append(args, payment.CountPay, payment.Id, payment.ChatId)
	values := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE auto_payments SET %s WHERE id = $%d and user_id = $%d", values, numParam, numParam+1)

	_, err := r.db.Exec(query, args...)

	return err
}
