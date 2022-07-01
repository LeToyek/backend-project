package repository

import (
	"coin-batam/entities"
	"coin-batam/queries"
)

type User interface {
	AddUser(user entities.User) error
}

// GetUsers() ([]entities.User, error)
// GetUserById(ID string) (entities.User, error)
// Login(email string, password string) (string, error)
// }

func (r *Repository) AddUser(user entities.User) error {
	_, err := r.DB.Exec(
		queries.AddUserQ,
		user.User_id,
		user.First_name,
		user.Last_name,
		user.Email,
		user.Password,
		user.Phone,
		user.Created_at,
		user.Updated_at,
	)
	return err
}
