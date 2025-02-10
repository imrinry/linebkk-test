package account

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{}
}

func (m *MockAccountRepository) GetAccountByUserID(ctx context.Context, userID string, offset int, limit int) ([]Account, error) {
	rets := m.Called(ctx, userID, offset, limit)
	return rets.Get(0).([]Account), rets.Error(1)
}

func (m *MockAccountRepository) GetCountAccounts(ctx context.Context, userID string) (int, error) {
	rets := m.Called(ctx, userID)
	return rets.Get(0).(int), rets.Error(1)
}

func (m *MockAccountRepository) SetAccountCache(ctx context.Context, key string, value []AccountResponse, expiration time.Duration) error {
	rets := m.Called(ctx, key, value, expiration)
	return rets.Error(0)
}

func (m *MockAccountRepository) GetAccountCache(ctx context.Context, key string) ([]AccountResponse, error) {
	rets := m.Called(ctx, key)
	return rets.Get(0).([]AccountResponse), rets.Error(1)
}
