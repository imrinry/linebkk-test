package banner_test

import (
	"context"
	"errors"
	"line-bk-api/internal/banner"
	"line-bk-api/pkg/utils"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBannerByUserID(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		page          int
		limit         int
		cacheKey      string
		cacheHit      bool
		cachedData    []banner.BannerResponse
		mockBanner    []banner.Banner
		mockBannerErr error
		mockTotal     int
		mockTotalErr  error
		expectedTotal int
		expectedErr   string
	}{
		{
			name:     "success get banner from db",
			userID:   "user1",
			page:     1,
			limit:    10,
			cacheKey: "banner:user1:0:10",
			cacheHit: false,
			mockBanner: []banner.Banner{
				{
					BannerID:    "1",
					UserID:      "user1",
					Title:       "Banner 1",
					Description: "Description 1",
					Image:       "image1.jpg",
				},
			},
			mockTotal:     1,
			expectedTotal: 1,
		},
		{
			name:     "success get banner from cache",
			userID:   "user1",
			page:     1,
			limit:    10,
			cacheKey: "banner:user1:0:10",
			cacheHit: true,
			cachedData: []banner.BannerResponse{
				{
					BannerID:    "1",
					Title:       "Banner 1",
					Description: "Description 1",
					Image:       "image1.jpg",
				},
			},
			expectedTotal: 0,
		},
		{
			name:          "error getting banner from db",
			userID:        "user1",
			page:          1,
			limit:         10,
			cacheKey:      "banner:user1:0:10",
			cacheHit:      false,
			mockBannerErr: errors.New("db error"),
			expectedErr:   "db error",
		},
		{
			name:     "error getting total banner count",
			userID:   "user1",
			page:     1,
			limit:    10,
			cacheKey: "banner:user1:0:10",
			cacheHit: false,
			mockBanner: []banner.Banner{
				{
					BannerID:    "1",
					UserID:      "user1",
					Title:       "Banner 1",
					Description: "Description 1",
					Image:       "image1.jpg",
				},
			},
			mockTotalErr: errors.New("total count error"),
			expectedErr:  "total count error",
		},
		{
			name:     "error getting banner from cache",
			userID:   "user1",
			page:     1,
			limit:    10,
			cacheKey: "banner:user1:0:10",
			cacheHit: false,
			mockBanner: []banner.Banner{
				{
					BannerID:    "1",
					UserID:      "user1",
					Title:       "Banner 1",
					Description: "Description 1",
					Image:       "image1.jpg",
				},
			},
			mockTotal:     1,
			expectedTotal: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := banner.NewMockBannerRepository()
			offset, limit := utils.GetOffset(tt.page, tt.limit)

			if !tt.cacheHit {
				mockRepo.On("GetBannerCache", context.Background(), tt.cacheKey).Return([]banner.BannerResponse{}, redis.Nil)
				mockRepo.On("GetBannerByUserID", context.Background(), tt.userID, offset, limit).Return(tt.mockBanner, tt.mockBannerErr)
				if tt.mockBannerErr == nil {
					mockRepo.On("GetTotalBannerByUserID", context.Background(), tt.userID).Return(tt.mockTotal, tt.mockTotalErr)
					if tt.mockTotalErr == nil {
						mockRepo.On("SetBannerCache", context.Background(), tt.cacheKey, mock.Anything, 5*time.Minute).Return(nil)
					}
				}
			} else {
				mockRepo.On("GetBannerCache", context.Background(), tt.cacheKey).Return(tt.cachedData, nil)
			}

			service := banner.NewBannerService(mockRepo)

			// Act
			banners, total, err := service.GetBannerByUserID(context.Background(), tt.userID, tt.page, tt.limit)

			// Assert
			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedTotal, total)

			if tt.cacheHit {
				assert.Equal(t, tt.cachedData, banners)
			} else if len(banners) > 0 {
				expectedResponses := make([]banner.BannerResponse, len(tt.mockBanner))
				for i, b := range tt.mockBanner {
					expectedResponses[i] = b.ToBannerResponse()
				}
				assert.Equal(t, expectedResponses, banners)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestNewBannerService(t *testing.T) {
	// Arrange
	mockRepo := banner.NewMockBannerRepository()

	// Act
	service := banner.NewBannerService(mockRepo)

	// Assert
	assert.NotNil(t, service)
	assert.Implements(t, (*banner.BannerService)(nil), service)
}
