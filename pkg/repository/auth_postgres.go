package repository

import (
	"database/sql"
	"fmt"

	TonWork "github.com/TonWork/back"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user TonWork.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, subscribe, telefon, position, description, companies, name, surname) VALUES ($1,$2,$3,'free','-','-','-','-','-','-') RETURNING id", Table_users)
	row := tx.QueryRow(query, user.Username, user.Password_hash, user.Email)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	query2 := fmt.Sprintf("INSERT INTO %s (id_user, time_in_hours_to_end) VALUES ($1, $2)", Table_user_sub)
	_, err = tx.Exec(query2, id, -1)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *AuthPostgres) GetUser(username, password string) (TonWork.User, error) {
	var user TonWork.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", Table_users)
	err := r.db.QueryRow(query, username, password).Scan(&user)
	return user, err
}
func (r *AuthPostgres) GetUserPorfile(username string) (TonWork.User, error) {
	var user TonWork.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", Table_users)
	err := r.db.QueryRow(query, username).Scan(&user)
	return user, err
}
