package banner

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type BannerRepository interface {
	GetBannerByUserID(ctx context.Context, userID string, offset int, limit int) ([]Banner, error)
	GetTotalBannerByUserID(ctx context.Context, userID string) (int, error)
}

type bannerRepository struct {
	db *sqlx.DB
}

func NewBannerRepository(db *sqlx.DB) BannerRepository {
	return &bannerRepository{db: db}
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
