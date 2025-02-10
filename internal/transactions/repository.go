package transactions

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	GetTransactionByUserID(ctx context.Context, userID string, offset int, limit int) ([]Transaction, error)
	GetTransactionCountByUserID(ctx context.Context, userID string) (int, error)
	SetTransactionCache(ctx context.Context, key string, value []TransactionResponse, expiration time.Duration) error
	GetTransactionCache(ctx context.Context, key string) ([]TransactionResponse, error)
}

type transactionRepository struct {
	db    *sqlx.DB
	cache *redis.Client
}

func NewTransactionRepository(db *sqlx.DB, cache *redis.Client) TransactionRepository {
	return &transactionRepository{db: db, cache: cache}
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

func (r *transactionRepository) SetTransactionCache(ctx context.Context, key string, value []TransactionResponse, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, key, jsonValue, expiration).Err()
}

func (r *transactionRepository) GetTransactionCache(ctx context.Context, key string) ([]TransactionResponse, error) {
	value, err := r.cache.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var transactions []TransactionResponse
	err = json.Unmarshal([]byte(value), &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
