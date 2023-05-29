package service

import (
	repository "github.com/TonWork/back/pkg/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
