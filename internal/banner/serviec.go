package banner

import (
	"context"
	"errors"
	"fmt"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"

	"time"

	"github.com/go-redis/redis/v8"
)

type BannerService interface {
	GetBannerByUserID(ctx context.Context, userID string, page int, limit int) ([]BannerResponse, int, error)
}

type bannerService struct {
	bannerRepository BannerRepository
}

func NewBannerService(bannerRepository BannerRepository) BannerService {
	return &bannerService{
		bannerRepository: bannerRepository,
	}
}

func (s *bannerService) GetBannerByUserID(ctx context.Context, userID string, page int, limit int) ([]BannerResponse, int, error) {
	offset, limit := utils.GetOffset(page, limit)
	cacheKey := fmt.Sprintf("banner:%s:%d:%d", userID, offset, limit)
	bannerResponses, err := s.bannerRepository.GetBannerCache(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		logs.Error(err)
		return []BannerResponse{}, 0, err
	}

	if len(bannerResponses) > 0 {
		return bannerResponses, 0, nil
	}

	banner, err := s.bannerRepository.GetBannerByUserID(ctx, userID, offset, limit)
	if err != nil {
		logs.Error(err)
		return []BannerResponse{}, 0, err
	}
	bannerResponses = make([]BannerResponse, len(banner))
	for i, b := range banner {
		bannerResponses[i] = b.ToBannerResponse()
	}

	total, err := s.bannerRepository.GetTotalBannerByUserID(ctx, userID)
	if err != nil {
		return bannerResponses, 0, err
	}

	// set redis cache
	err = s.bannerRepository.SetBannerCache(ctx, cacheKey, bannerResponses, 5*time.Minute)
	if err != nil {
		logs.Error(err)
	}

	return bannerResponses, total, nil
}
