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
func (r *SubscribesPostgres) CancelSubscribe(id int) error {
	query := fmt.Sprintf("UPDATE %s SET subscribe='free' WHERE id=$1", Table_users)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SubscribesPostgres) ChangeSubscribeTime() error {
	query := fmt.Sprintf("UPDATE %s SET time_in_hours_to_end=time_in_hours_to_end-1 WHERE time_in_hours_to_end<>-1, subscribe='premium'", Table_users)
	query2 := fmt.Sprintf("UPDATE %s SET subscribe='free' WHERE time_in_hours_to_end<=-1", Table_users)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query2)
	if err != nil {
		return err
	}
	return nil
}
