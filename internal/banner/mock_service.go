package banner

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockBannerService struct {
	mock.Mock
}

func NewMockBannerService() *MockBannerService {
	return &MockBannerService{}
}

func (m *MockBannerService) GetBannerByUserID(ctx context.Context, userID string, page int, limit int) ([]BannerResponse, int, error) {
	args := m.Called(ctx, userID, page, limit)
	return args.Get(0).([]BannerResponse), args.Get(1).(int), args.Error(2)
}
