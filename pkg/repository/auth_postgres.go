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
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, subscribe, telefon, position, description, companies, name, surname) OUTPUT inserted.id VALUES (@Username,@Password_hash,@Email,'free','-','-','-','-','-','-')", Table_users)
	row := tx.QueryRow(query, sql.Named("Username", user.Username), sql.Named("Password_hash", user.Password_hash), sql.Named("Email", user.Email))
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
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=@Username AND password_hash=@Password", Table_users)
	err := r.db.QueryRow(query, sql.Named("Username", username), sql.Named("Password", password)).Scan(&user.Id, &user.Username, &user.Password_hash, &user.Email, &user.Telefon, &user.Position, &user.Description, &user.Subscribe, &user.Companies, &user.Name, &user.Surname)
	return user, err
}
func (r *AuthPostgres) GetUserPorfile(username string) (TonWork.User, error) {
	var user TonWork.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=@Username", Table_users)
	err := r.db.QueryRow(query, sql.Named("Username", username)).Scan(&user.Id, &user.Username, &user.Password_hash, &user.Email, &user.Telefon, &user.Position, &user.Description, &user.Subscribe, &user.Companies, &user.Name, &user.Surname)
	return user, err
}
