package postgres

import (
	"AutoPayment/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDb(cfg config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.URI())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
