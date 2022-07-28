package entities

type User struct {
	User_id    string `json:"user_id"`
	First_name string `json:"first_name" binding:"required,min=2,max=100"`
	Last_name  string `json:"last_name" binding:"required,min=2,max=100"`
	Email      string `json:"email" binding:"email,required"`
	Password   string `json:"password" binding:"required,min=6"`
	Phone      string `json:"phone" binding:"required"`
	Created_at string `json:"Created_at"`
	Updated_at string `json:"Updated_at"`
}
