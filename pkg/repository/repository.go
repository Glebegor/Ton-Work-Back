package repository

import (
	TonWork "github.com/TonWork/back"
	sqlx "github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(TonWork.User) error
	GetUser(username, password string) (TonWork.User, error)
	GetUserPorfile(username string) (TonWork.User, error)
}
type Work interface {
	GetAll() ([]TonWork.Work, error)
}
type Posts interface {
}
type Subscribes interface {
}
type Repository struct {
	Authorization
	Work
	Posts
	Subscribes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Work:          NewWorkPostgres(db),
		Posts:         nil,
		Subscribes:    nil,
	}
}
