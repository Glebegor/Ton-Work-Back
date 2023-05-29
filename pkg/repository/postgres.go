package repository

import (
	"fmt"

	sqlx "github.com/jmoiron/sqlx"
)

type ConfigDB struct {
	Host     string
	Port     string
	DBName   string
	SSLMode  string
	Username string
	Password string
}

func ConnectDB(cfg ConfigDB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
