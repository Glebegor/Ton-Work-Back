package service

import (
	TonWork "github.com/TonWork/back"
	repository "github.com/TonWork/back/pkg/repository"
)

type Authorization interface {
	CreateUser(user TonWork.User) error
	PasswordHash(password string) string
	GenerateToken(username, password string) (string, error)
	ParseToken(accesToken string) (int, string, string, string, error)
	GetUserProfile(param string) (TonWork.User, error)
}
type Work interface {
	GetAll([]TonWork.Work, error)
	Create(int, TonWork.Work) error
}
type Posts interface {
}
type Subscribes interface {
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
		Posts:         nil,
		Subscribes:    nil,
	}
}
