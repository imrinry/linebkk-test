package banner

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockBannerRepository struct {
	mock.Mock
}

func NewMockBannerRepository() *MockBannerRepository {
	return &MockBannerRepository{}
}

func (m *MockBannerRepository) GetBannerByUserID(ctx context.Context, userID string, offset int, limit int) ([]Banner, error) {
	args := m.Called(ctx, userID, offset, limit)
	return args.Get(0).([]Banner), args.Error(1)
}

func (m *MockBannerRepository) GetTotalBannerByUserID(ctx context.Context, userID string) (int, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int), args.Error(1)
}
