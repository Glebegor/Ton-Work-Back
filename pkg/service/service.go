package service

import (
	repository "github.com/TonWork/back/pkg/repository"
)

type Authorization interface {
}
type Work interface {
}
type Posts interface {
}
type Subscribes interface {
}

type Service struct {
	repo repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
