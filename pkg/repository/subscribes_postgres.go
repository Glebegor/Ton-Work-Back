package repository

import (
	"database/sql"
	"fmt"
)

type SubscribesPostgres struct {
	db *sql.DB
}
type SubTime struct {
	Time_in_hours_to_end int `json:"time_in_hours_to_end" db:"time_in_hours_to_end"`
	UserId               int `json:"userId" db:"id_user"`
	Id                   int `json:"id" db:"id"`
}

func NewSubscribesPostgres(db *sql.DB) *SubscribesPostgres {
	return &SubscribesPostgres{db: db}
}
func (r *SubscribesPostgres) BuySubscribe(id int) error {
	query := fmt.Sprintf("UPDATE %s SET subscribe='premium' WHERE id=$1", Table_users)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	query2 := fmt.Sprintf("UPDATE %s SET time_in_hours_to_end=30 WHERE id_user=$1", Table_user_sub)
	_, err = r.db.Exec(query2, id)
	return err
}
func (r *SubscribesPostgres) CancelSubscribe(id int) error {
	query := fmt.Sprintf("UPDATE %s SET subscribe='free' WHERE id=$1", Table_users)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	query2 := fmt.Sprintf("UPDATE %s SET time_in_hours_to_end=-1 WHERE id_user=$1", Table_user_sub)
	_, err = r.db.Exec(query2, id)
	return err
}
func (r *SubscribesPostgres) GetTimeToEnd(id int) (int, error) {
	var SubTimeInfo SubTime
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_user=@Id", Table_user_sub)
	err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(&SubTimeInfo.Id, &SubTimeInfo.Time_in_hours_to_end, &SubTimeInfo.UserId)
	return SubTimeInfo.Time_in_hours_to_end, err
}
func (r *SubscribesPostgres) UpdateTimeOfSub() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE tl SET tl.time_in_hours_to_end=tl.time_in_hours_to_end-1 FROM %s tl JOIN %s ul ON tl.id_user = ul.id WHERE tl.time_in_hours_to_end<>-1 AND ul.subscribe='premium'", Table_user_sub, Table_users)
	query2 := fmt.Sprintf("UPDATE tl SET subscribe=SUBSTRING('free', 1, 255) FROM %s tl JOIN %s ul ON tl.id=ul.id_user WHERE ul.time_in_hours_to_end<=-1", Table_users, Table_user_sub)

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query2)
	if err != nil {
		return err
	}
	return tx.Commit()
}
