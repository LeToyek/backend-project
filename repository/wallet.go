package repository

type Wallet interface {
	GetBalance(userID string)
}
