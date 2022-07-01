package service

import "coin-batam/entities"

type User interface {
	AddUser(user entities.User) error
}

func (s *Service) AddUser(user entities.User) error {
	return s.Repository.AddUser(user)
}
