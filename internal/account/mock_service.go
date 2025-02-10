package account

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func NewMockAccountService() *MockAccountService {
	return &MockAccountService{}
}

func (m *MockAccountService) GetAccountByUserID(ctx context.Context, userID string, page int, limit int) ([]AccountResponse, int, error) {
	rets := m.Called(ctx, userID, page, limit)
	return rets.Get(0).([]AccountResponse), rets.Get(1).(int), rets.Error(2)
}
