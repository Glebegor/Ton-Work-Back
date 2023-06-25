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
func (s *SubscribesService) CancelSubscribe(id int) error {
	err := s.repos.CancelSubscribe(id)
	return err
}
func (s *SubscribesService) GetTimeToEnd(id int) (int, error) {
	timetoend, err := s.repos.GetTimeToEnd(id)
	return timetoend, err
}
