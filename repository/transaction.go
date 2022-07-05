package repository

import (
	"coin-batam/entities"
	"coin-batam/queries"
	"log"
)

type Transaction interface {
	AddTransaction(transaction entities.Transaction) error
	GetTransaction(userID string) ([]entities.Transaction, error)
}

func (r *Repository) AddTransaction(transaction entities.Transaction) error {
	_, err := r.DB.Exec(
		queries.AddTransactionQ,
		transaction.Transaction_id,
		transaction.Activity,
		transaction.Coin_ID,
		transaction.Price,
		transaction.User_id,
		transaction.Created_at,
		transaction.Updated_at,
	)
	return err
}
func (r *Repository) GetTransaction(userID string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	rows, err := r.DB.Query(
		queries.GetTransactionQ,
		userID,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		transaction := entities.Transaction{}
		err := rows.Scan(
			&transaction.Transaction_id,
			&transaction.Activity,
			&transaction.Coin_ID,
			&transaction.Price,
			&transaction.User_id,
			&transaction.Created_at,
			&transaction.Updated_at,
		)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err
}
