package banner

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBannerByUserID(t *testing.T) {

	// Arrange
	mockRepo := NewMockBannerRepository()

	mockRepo.On("GetBannerByUserID", context.Background(), "1", 0, 10).Return([]Banner{
		{
			BannerID:    "1",
			UserID:      "1",
			Title:       "Banner 1",
			Description: "Description 1",
			Image:       "Image 1",
		},
	}, nil)

	mockRepo.On("GetTotalBannerByUserID", context.Background(), "1").Return(1, nil)

	service := NewBannerService(mockRepo)

	banners, total, err := service.GetBannerByUserID(context.Background(), "1", 0, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, 1, total)
	assert.Equal(t, "1", banners[0].BannerID)
	assert.Equal(t, "Banner 1", banners[0].Title)
	assert.Equal(t, "Description 1", banners[0].Description)
	assert.Equal(t, "Image 1", banners[0].Image)

}

func TestGetBannerByUserID_Error(t *testing.T) {
	// Arrange
	mockRepo := NewMockBannerRepository()

	mockRepo.On("GetBannerByUserID", context.Background(), "1", 0, 10).Return([]Banner{}, errors.New("error"))
	mockRepo.On("GetTotalBannerByUserID", context.Background(), "1").Return(0, errors.New("error"))

	service := NewBannerService(mockRepo)

	banners, total, err := service.GetBannerByUserID(context.Background(), "1", 0, 10)

	assert.Empty(t, banners)
	assert.Equal(t, 0, total)
	assert.Equal(t, "error", err.Error())
}

func TestGetBannerByUserID_CountError(t *testing.T) {
	// Arrange
	mockRepo := NewMockBannerRepository()

	mockRepo.On("GetBannerByUserID", context.Background(), "1", 0, 10).Return([]Banner{
		{
			BannerID:    "1",
			UserID:      "1",
			Title:       "Banner 1",
			Description: "Description 1",
			Image:       "Image 1",
		},
	}, nil)

	mockRepo.On("GetTotalBannerByUserID", context.Background(), "1").Return(0, errors.New("error"))

	service := NewBannerService(mockRepo)

	_, _, err := service.GetBannerByUserID(context.Background(), "1", 0, 10)

	// assert.Empty(t, banners)
	// assert.Equal(t, 0, total)
	assert.Equal(t, "error", err.Error())
}
