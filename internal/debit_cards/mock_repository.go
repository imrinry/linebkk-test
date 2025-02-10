package debit_cards

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockDebitCardRepository struct {
	mock.Mock
}

func (m *MockDebitCardRepository) GetDebitCardByUserID(ctx context.Context, userID string, offset int, limit int) ([]DebitCard, error) {
	rets := m.Called(ctx, userID, offset, limit)
	return rets.Get(0).([]DebitCard), rets.Error(1)
}

func (m *MockDebitCardRepository) GetCountDebitCards(ctx context.Context, userID string) (int, error) {
	rets := m.Called(ctx, userID)
	return rets.Get(0).(int), rets.Error(1)
}

func (m *MockDebitCardRepository) SetDebitCardCache(ctx context.Context, key string, value []DebitCardResponse, expiration time.Duration) error {
	rets := m.Called(ctx, key, value, expiration)
	return rets.Error(0)
}

func (m *MockDebitCardRepository) GetDebitCardCache(ctx context.Context, key string) ([]DebitCardResponse, error) {
	rets := m.Called(ctx, key)
	return rets.Get(0).([]DebitCardResponse), rets.Error(1)
}
