package service

import "github.com/TonWork/back/pkg/repository"

type SubscribesService struct {
	repos repository.Subscribes
}

func NewSubscribesService(repo repository.Subscribes) *SubscribesService {
	return &SubscribesService{repos: repo}
}
func (s *SubscribesService) BuySubscribe(id int) error {
	err := s.repos.BuySubscribe(id)
	return err
}
