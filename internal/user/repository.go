package user

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetAllUsers(offset, limit int) ([]User, error)
}

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) GetAllUsers(offset, limit int) ([]User, error) {
	rows, err := r.db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
