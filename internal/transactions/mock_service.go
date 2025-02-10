package transactions

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockTransactionsService struct {
	mock.Mock
}

func NewMockTransactionsService() *MockTransactionsService {
	return &MockTransactionsService{}
}

func (m *MockTransactionsService) GetTransactionByUserID(ctx context.Context, userID string, page int, limit int) ([]TransactionResponse, int, error) {
	args := m.Called(ctx, userID, page, limit)
	return args.Get(0).([]TransactionResponse), args.Get(1).(int), args.Error(2)
}
