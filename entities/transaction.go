package entities

type Transaction struct {
	Transaction_id string  `json:"transaction_id"`
	Activity       string  `json:"activity"`
	Coin_ID        string  `json:"coin_id"`
	Price          float64 `json:"price"`
	User_id        string  `json:"user_id"`
	Created_at     string  `json:"created_at"`
	Updated_at     string  `json:"updated_at"`
}
