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
	Create(int, TonWork.Work) error
}
type Post interface {
	GetAll() ([]TonWork.Post, error)
	Create(int, TonWork.Post) error
	GetById(id int) (TonWork.Post, error)
}
type Subscribes interface {
}
type Repository struct {
	Authorization
	Work
	Post
	Subscribes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Work:          NewWorkPostgres(db),
		Post:          NewPostPostgres(db),
		Subscribes:    nil,
	}
}
