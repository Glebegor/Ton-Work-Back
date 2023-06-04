package repository

import (
	"fmt"

	TonWork "github.com/TonWork/back"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user TonWork.User) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, subscribe, telefon, position, description, companies, name, surname) VALUES ($1,$2,$3,'free','-','-','-','-','-','-') RETURNING id", Table_users)
	row := r.db.QueryRow(query, user.Username, user.Password_hash, user.Email)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}
func (r *AuthPostgres) GetUser(username, password string) (TonWork.User, error) {
	var user TonWork.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", Table_users)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
func (r *AuthPostgres) GetUserPorfile(username string) (TonWork.User, error) {
	var user TonWork.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", Table_users)
	err := r.db.Get(&user, query, username)
	return user, err
}
