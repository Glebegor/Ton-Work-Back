package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SubscribesPostgres struct {
	db *sqlx.DB
}
type SubTime struct {
	Time_in_hours_to_end int `json:"time_in_hours_to_end"`
	UserId               int `json:"userId"`
	Id                   int `json:"id"`
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
func (r *SubscribesPostgres) GetTimeToEnd(id int) (int, error) {
	var SubTimeInfo SubTime
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", Table_user_sub)
	err := r.db.Select(&SubTimeInfo, query, id)
	return SubTimeInfo.Time_in_hours_to_end, err
}

func (r *SubscribesPostgres) ChangeSubscribeTime() error {
	query := fmt.Sprintf("UPDATE %s tl SET time_in_hours_to_end=time_in_hours_to_end-1 FROM %s ul WHERE tl.time_in_hours_to_end<>-1 AND ul.subscribe='premium'", Table_user_sub, Table_users)
	query2 := fmt.Sprintf("UPDATE %s tl SET subscribe='free' FROM %s ul WHERE ul.time_in_hours_to_end<=-1", Table_users, Table_user_sub)
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
