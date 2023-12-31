package service

import (
	repository "github.com/Glebegor/Ton-Work-Back/pkg/repository"
	TonWork "github.com/Glebegor/Ton-Work-Back/structint"
)

type Authorization interface {
	CreateUser(user TonWork.User) error
	PasswordHash(password string) string
	GenerateToken(username, password string) (string, error)
	ParseToken(accesToken string) (int, string, string, string, error)
	GetUserProfile(param string) (TonWork.User, error)
}
type Work interface {
	GetAll() ([]TonWork.Work, error)
	Create(int, TonWork.Work) error
	GetById(int) (TonWork.Work, error)
	Update(int, TonWork.WorkUpdate) error
	Delete(int) error
}
type Posts interface {
	GetAll() ([]TonWork.Post, error)
	Create(int, TonWork.Post) error
	GetById(int) (TonWork.Post, error)
	Update(int, TonWork.PostUpdate) error
	Delete(int) error
}
type Subscribes interface {
	BuySubscribe(int) error
	CancelSubscribe(int) error
	GetTimeToEnd(int) (int, error)
}

type Service struct {
	Authorization
	Work
	Posts
	Subscribes
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Work:          NewWorkService(repos.Work),
		Posts:         NewPostService(repos.Post),
		Subscribes:    NewSubscribesService(repos.Subscribes),
	}
}
