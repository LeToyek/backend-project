package service

import "coin-batam/entities"

type Transaction interface {
	AddTransaction(transaction entities.Transaction) error
	GetTransaction(userID string) ([]entities.Transaction, error)
}

func (s *Service) AddTransaction(transaction entities.Transaction) error {
	return s.Repository.AddTransaction(transaction)
}

func (s *Service) GetTransaction(userID string) ([]entities.Transaction, error) {
	return s.Repository.GetTransaction(userID)
}
