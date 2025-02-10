package auth

import "github.com/jmoiron/sqlx"

type AuthRepository interface {
	LoginWithPinCode(userID string, pinCode string) error
	LoginWithPassword(userID string, password string) error
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) LoginWithPinCode(userID string, pinCode string) error {
	return nil
}

func (r *authRepository) LoginWithPassword(userID string, password string) error {
	return nil
}
