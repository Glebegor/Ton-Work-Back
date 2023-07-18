package repository

import (
	"database/sql"
	"fmt"
)

const (
	Table_users          = "users"
	Table_works          = "works"
	Table_posts          = "posts"
	Table_user_sub       = "user_sub"
	Table_users_to_works = "users_to_works"
	Table_users_to_posts = "users_to_posts"
)

type ConfigDB struct {
	Server   string
	User     string
	Password string
	Port     string
	Database string
}

func ConnectDB(cfg ConfigDB) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", cfg.Server, cfg.User, cfg.Password, cfg.Port, cfg.Database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
