package user

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetAllUsers(offset, limit int) ([]User, error)
	GetCountUsers() (int, error)
	GetUserByID(id string) (User, error)
	GetUserGreeting(id string, offset, limit int) ([]UserGreeting, error)
	GetUserGreetingCount(id string) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) GetAllUsers(offset, limit int) ([]User, error) {
	rows, err := r.db.Queryx(`
		SELECT 
			user_id, 
			name, 
			dummy_col_1 
		FROM 
			users 
		ORDER BY 
			user_id ASC
		LIMIT ?, ?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Dummy); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *repository) GetCountUsers() (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users")
	return count, err
}

func (r *repository) GetUserByID(id string) (User, error) {
	var user User
	err := r.db.QueryRowx("SELECT user_id, name, dummy_col_1, email, phone_number, profile_image, pin_code, password, created_at FROM users WHERE user_id = ?", id).StructScan(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) GetUserGreeting(id string, offset, limit int) ([]UserGreeting, error) {
	var greetings []UserGreeting
	rows, err := r.db.Queryx(`
		SELECT 
			user_id, 
			greeting, 
			dummy_col_2 
		FROM 
			user_greetings 
		WHERE 
			user_id = ? 
		ORDER BY 
			greeting DESC 
		LIMIT ?, ?`, id, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var greeting UserGreeting
		if err := rows.Scan(&greeting.UserID, &greeting.Greeting, &greeting.Dummy); err != nil {
			return nil, err
		}
		greetings = append(greetings, greeting)
	}
	return greetings, nil
}

func (r *repository) GetUserGreetingCount(id string) (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM user_greetings WHERE user_id = ?", id)
	return count, err
}
