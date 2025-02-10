package transactions

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	GetTransactionByUserID(ctx context.Context, userID string, offset int, limit int) ([]Transaction, error)
	GetTransactionCountByUserID(ctx context.Context, userID string) (int, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetTransactionByUserID(ctx context.Context, userID string, offset int, limit int) ([]Transaction, error) {
	query := `SELECT transaction_id, user_id, name, image, isBank 
			  FROM 
			  		transactions 
			  WHERE 
			  		user_id = ? 
			  ORDER BY 
			  		transaction_id ASC 
			  LIMIT ? OFFSET ?`
	var transactions []Transaction
	err := r.db.Select(&transactions, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransactionCountByUserID(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) FROM transactions WHERE user_id = ?`
	var count int
	err := r.db.GetContext(ctx, &count, query, userID)
	return count, err
}
