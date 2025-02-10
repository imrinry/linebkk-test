package debit_cards

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDebitCardService struct {
	mock.Mock
}

func (m *MockDebitCardService) GetDebitCards(ctx context.Context, userID string, page int, limit int) ([]DebitCardResponse, int, error) {
	rets := m.Called(ctx, userID, page, limit)
	return rets.Get(0).([]DebitCardResponse), rets.Get(1).(int), rets.Error(2)
}
