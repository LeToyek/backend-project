package entities

type Wallet struct {
	Wallet_id string  `json:"wallet_id"`
	Balance   float64 `json:"balance"`
	User_id   string  `json:"user_id"`
}
