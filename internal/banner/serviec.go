package banner

import (
	"context"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
)

type BannerService interface {
	GetBannerByUserID(ctx context.Context, userID string, page int, limit int) ([]BannerResponse, int, error)
}

type bannerService struct {
	bannerRepository BannerRepository
}

func NewBannerService(bannerRepository BannerRepository) BannerService {
	return &bannerService{bannerRepository: bannerRepository}
}

func (s *bannerService) GetBannerByUserID(ctx context.Context, userID string, page int, limit int) ([]BannerResponse, int, error) {

	offset, limit := utils.GetOffset(page, limit)
	banner, err := s.bannerRepository.GetBannerByUserID(ctx, userID, offset, limit)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	bannerResponses := make([]BannerResponse, len(banner))
	for i, b := range banner {
		bannerResponses[i] = b.ToBannerResponse()
	}

	total, err := s.bannerRepository.GetTotalBannerByUserID(ctx, userID)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	return bannerResponses, total, nil
}
