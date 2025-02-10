package transactions

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockTransactionsRepository struct {
	mock.Mock
}

func NewMockTransactionsRepository() *MockTransactionsRepository {
	return &MockTransactionsRepository{}
}

func (m *MockTransactionsRepository) GetTransactionByUserID(ctx context.Context, userID string, offset int, limit int) ([]Transaction, error) {
	args := m.Called(ctx, userID, offset, limit)
	return args.Get(0).([]Transaction), args.Error(1)
}

func (m *MockTransactionsRepository) GetTransactionCountByUserID(ctx context.Context, userID string) (int, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockTransactionsRepository) SetTransactionCache(ctx context.Context, key string, value []TransactionResponse, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockTransactionsRepository) GetTransactionCache(ctx context.Context, key string) ([]TransactionResponse, error) {
	args := m.Called(ctx, key)
	return args.Get(0).([]TransactionResponse), args.Error(1)
}
