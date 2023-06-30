package postgres

import (
	"AutoPayment/internal/model"
	"database/sql"
	"errors"
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

func (r *PaymentRepository) IndexByTime(limit, offset int, time time.Time) ([]model.Payment, error) {
	var payments []model.Payment

	query := "SELECT * FROM auto_payments WHERE next_pay_date < $1 LIMIT $2 OFFSET $3"

	err := r.db.Select(&payments, query, time, limit, offset)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *PaymentRepository) Create(payment model.Payment) (model.Payment, error) {
	query := `INSERT INTO auto_payments (chat_id, name, period_day, payment_day, amount, count_pay, next_pay_date, created_at) 
			  VALUES (:chat_id, :name, :period_day, :payment_day, :amount, :count_pay, :next_pay_date, :created_at) 
			  RETURNING *`
	payment.CreatedAt = time.Now()

	result, err := r.db.NamedQuery(query, payment)
	if err != nil {
		return model.Payment{}, err
	}
	defer result.Close()
	for result.Next() {
		err = result.StructScan(&payment)
		if err != nil {
			return model.Payment{}, err
		}
	}

	return payment, nil
}

func (r *PaymentRepository) ExistsByName(chatId int64, name string) (bool, error) {
	var payment model.Payment
	query := "SELECT * FROM auto_payments WHERE chat_id = $1 and name = $2"

	err := r.db.Get(&payment, query, chatId, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PaymentRepository) GetByName(chatId int64, name string) (model.Payment, error) {
	var payment model.Payment
	query := "SELECT * FROM auto_payments WHERE chat_id = $1 and name = $2"

	err := r.db.Get(&payment, query, chatId, name)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *PaymentRepository) IndexByChatId(chatId int64) ([]model.Payment, error) {
	var models []model.Payment
	query := "SELECT * FROM auto_payments WHERE chat_id = $1"

	err := r.db.Select(&models, query, chatId)

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (r *PaymentRepository) Show(id int) (model.Payment, error) {
	payment := model.Payment{}
	query := "SELECT * FROM auto_payments WHERE id = $1"

	err := r.db.Get(&payment, query, id)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *PaymentRepository) Delete(chatId int64, name string) error {
	query := "DELETE FROM auto_payments WHERE chat_id = $1 and name = $2"
	res, err := r.db.Exec(query, chatId, name)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no one rows was delete")
	}

	return err
}

func (r *PaymentRepository) Update(updatePayment model.UpdatePayment) (model.Payment, error) {
	args := make([]interface{}, 0)
	setValues := make([]string, 0)
	numParam := 1

	if updatePayment.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", numParam))
		args = append(args, updatePayment.Name)
		numParam++
	}

	if updatePayment.PeriodDay != nil {
		setValues = append(setValues, fmt.Sprintf("period_day = $%d", numParam))
		args = append(args, updatePayment.PeriodDay)
		numParam++
	}

	if updatePayment.PaymentDay != nil {
		setValues = append(setValues, fmt.Sprintf("payment_day = $%d", numParam))
		args = append(args, updatePayment.PaymentDay)
		numParam++
	}

	if updatePayment.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount = $%d", numParam))
		args = append(args, updatePayment.Amount)
		numParam++
	}

	if updatePayment.NextPayDate != nil {
		setValues = append(setValues, fmt.Sprintf("next_pay_date = $%d", numParam))
		args = append(args, updatePayment.NextPayDate)
		numParam++
	}

	if updatePayment.CountPay != nil {
		setValues = append(setValues, fmt.Sprintf("count_pay = $%d", numParam))
		args = append(args, updatePayment.CountPay)
		numParam++
	}

	args = append(args, updatePayment.ID)
	values := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE auto_payments SET %s WHERE id = $%d RETURNING *", values, numParam)
	payment := model.Payment{}

	err := r.db.Get(&payment, query, args...)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *PaymentRepository) SumForMonth(chatId int64) (int, error) {
	query := "SELECT sum(amount * (30 / period_day)) FROM auto_payments WHERE chat_id = $1"
	var sum int

	err := r.db.Get(&sum, query, chatId)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
