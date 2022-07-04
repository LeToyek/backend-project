package repository

import (
	"coin-batam/entities"
	"coin-batam/queries"
	"database/sql"
	"errors"
)

type User interface {
	AddUser(user entities.User) error
	UseUser(email string, password string) (string, error)
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
func (r *Repository) UseUser(email string, password string) (string, error) {
	var user entities.User
	row := r.DB.QueryRow(queries.FindUserQ, email)
	err := row.Scan(
		&user.Email,
		&user.Password,
		&user.User_id,
	)
	if err == sql.ErrNoRows {
		return "", errors.New("email not found")
	}
	if password != user.Password {
		return "", errors.New("wrong password")
	}
	if err != nil {
		return "", errors.New(err.Error())
	}
	return user.User_id, err
}
