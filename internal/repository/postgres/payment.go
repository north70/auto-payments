package postgres

import (
	"AutoPayment/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type paymentDto struct {
	Id         int    `db:"id"`
	UserId     int    `db:"user_id"`
	Name       string `db:"name"`
	PeriodType int    `db:"period_type"`
	PeriodDay  int    `db:"period_day"`
	PaymentDay int    `db:"payment_day"`
	Amount     int    `db:"amount"`
	CountPay   *int   `db:"count_pay"`
	CreatedAt  string `db:"created_at"`
}

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payment model.Payment) error {
	query := fmt.Sprintf("INSERT INTO auto_payments (chat_id, name, period_type, period_day, payment_day, amount, count_pay, created_at) VALUES (:user_id, :name, :period_type, :period_day, :payment_day, :amount, :count_pay, :created_at)")
	payment.CreatedAt = time.Now()
	dto := structToSql(payment)

	_, err := r.db.NamedExec(query, dto)

	return err
}

func (r *PaymentRepository) Index(userId int) ([]model.Payment, error) {
	var dtoModels []paymentDto
	query := fmt.Sprintf("SELECT * FROM auto_payments WHERE chat_id = $1")

	err := r.db.Select(&dtoModels, query, userId)

	if err != nil {
		return nil, err
	}

	var models []model.Payment
	for _, dtoModel := range dtoModels {
		trueModel, err := sqlToStruct(dtoModel)
		if err != nil {
			return nil, err
		}
		models = append(models, trueModel)
	}

	return models, nil
}

func (r *PaymentRepository) Show(userId, id int) (model.Payment, error) {
	dtoModel := paymentDto{}
	query := fmt.Sprintf("SELECT * FROM auto_payments WHERE id = $1 and chat_id = $2")

	err := r.db.Get(&dtoModel, query, id, userId)
	if err != nil {
		return model.Payment{}, err
	}

	payment, err := sqlToStruct(dtoModel)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *PaymentRepository) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM auto_payments WHERE chat_id = $1 and id = $2")
	_, err := r.db.Exec(query, userId, id)

	return err
}

func (r *PaymentRepository) Update(payment model.UpdatePayment) error {
	args := make([]interface{}, 0)
	setValues := make([]string, 0)
	numParam := 3

	if payment.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", numParam))
		args = append(args, payment.Name)
		numParam++
	}

	if payment.PeriodType != nil {
		setValues = append(setValues, fmt.Sprintf("period_type = $%d", numParam))
		args = append(args, payment.PeriodType)
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

	setValues = append(setValues, fmt.Sprintf("count_pay = $%d", numParam))

	args = append(args, payment.CountPay, payment.Id, payment.UserId)
	values := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE auto_payments SET %s WHERE id = $1 and user_id = $2", values)

	_, err := r.db.Exec(query, args...)

	return err
}

func sqlToStruct(dtoModel paymentDto) (model.Payment, error) {
	createdAt, err := time.Parse(time.RFC3339, dtoModel.CreatedAt)
	if err != nil {
		return model.Payment{}, err
	}

	return model.Payment{
		Id:         dtoModel.Id,
		ChatId:     dtoModel.UserId,
		Name:       dtoModel.Name,
		PeriodType: model.PeriodType(dtoModel.PeriodDay),
		PeriodDay:  dtoModel.PeriodDay,
		PaymentDay: dtoModel.PaymentDay,
		Amount:     dtoModel.Amount,
		CountPay:   dtoModel.CountPay,
		CreatedAt:  createdAt,
	}, nil
}

func structToSql(model model.Payment) paymentDto {

	return paymentDto{
		Id:         model.Id,
		UserId:     model.ChatId,
		Name:       model.Name,
		PeriodType: int(model.PeriodType),
		PeriodDay:  model.PeriodDay,
		PaymentDay: model.PaymentDay,
		Amount:     model.Amount,
		CountPay:   model.CountPay,
		CreatedAt:  model.CreatedAt.Format(time.RFC3339),
	}
}
