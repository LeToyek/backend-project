package service

import "coin-batam/entities"

type User interface {
	AddUser(user entities.User) error
	UseUser(email string, password string) (string, error)
}

func (s *Service) AddUser(user entities.User) error {
	return s.Repository.AddUser(user)
}
func (s *Service) UseUser(email string, password string) (string, error) {
	return s.Repository.UseUser(email, password)
}
