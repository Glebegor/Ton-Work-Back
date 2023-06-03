package service

import (
	TonWork "github.com/TonWork/back"
	repository "github.com/TonWork/back/pkg/repository"
)

type Authorization interface {
	CreateUser(user TonWork.User) error
	PasswordHash(password string) string
	GenerateToken(input TonWork.UserPerson) (string, error)
}
type Work interface {
}
type Posts interface {
}
type Subscribes interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
