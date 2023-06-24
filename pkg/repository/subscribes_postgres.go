package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SubscribesPostgres struct {
	db *sqlx.DB
}

func NewSubscribesPostgres(db *sqlx.DB) *SubscribesPostgres {
	return &SubscribesPostgres{db: db}
}
func (r *SubscribesPostgres) BuySubscribe(id int) error {
	query := fmt.Sprintf("UPDATE %s SET subscribe='premium' WHERE id=$1", Table_users)
	_, err := r.db.Exec(query, id)
	return err
}
