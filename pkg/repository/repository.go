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
	GetById(int) (TonWork.Work, error)
	Update(string, TonWork.WorkUpdate) error
	Delete(string) error
}
type Post interface {
	GetAll() ([]TonWork.Post, error)
	Create(int, TonWork.Post) error
	GetById(int) (TonWork.Post, error)
	Update(string, TonWork.PostUpdate) error
	Delete(string) error
}
type Subscribes interface {
	BuySubscribe(int) error
	CancelSubscribe(int) error
	ChangeSubscribeTime() error
	GetTimeToEnd(int) (int, error)
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
		Subscribes:    NewSubscribesPostgres(db),
	}
}
