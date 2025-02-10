//go:build !test

package banner

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type BannerRepository interface {
	GetBannerByUserID(ctx context.Context, userID string, offset int, limit int) ([]Banner, error)
	GetTotalBannerByUserID(ctx context.Context, userID string) (int, error)
	SetBannerCache(ctx context.Context, key string, value []BannerResponse, expiration time.Duration) error
	GetBannerCache(ctx context.Context, key string) ([]BannerResponse, error)
}

type bannerRepository struct {
	db    *sqlx.DB
	cache *redis.Client
}

func NewBannerRepository(db *sqlx.DB, cache *redis.Client) BannerRepository {
	return &bannerRepository{db: db, cache: cache}
}

func (r *bannerRepository) GetBannerByUserID(ctx context.Context, userID string, offset int, limit int) ([]Banner, error) {
	banner := []Banner{}
	query := `SELECT banner_id, title, description, image
	 		  FROM banners 
			  WHERE user_id = ? 
			  ORDER BY banner_id ASC 
			  LIMIT ? OFFSET ?`

	err := r.db.Select(&banner, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func (r *bannerRepository) GetTotalBannerByUserID(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) FROM banners WHERE user_id = ?`
	var total int
	err := r.db.GetContext(ctx, &total, query, userID)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *bannerRepository) SetBannerCache(ctx context.Context, key string, value []BannerResponse, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, key, jsonValue, expiration).Err()
}

func (r *bannerRepository) GetBannerCache(ctx context.Context, key string) ([]BannerResponse, error) {
	value, err := r.cache.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var banners []BannerResponse
	err = json.Unmarshal([]byte(value), &banners)
	if err != nil {
		return nil, err
	}
	return banners, nil
}
