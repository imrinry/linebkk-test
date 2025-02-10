package banner

type Banner struct {
	BannerID    string `db:"banner_id"`
	UserID      string `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Image       string `db:"image"`
}

func (b *Banner) ToBannerResponse() BannerResponse {
	return BannerResponse{
		BannerID:    b.BannerID,
		Title:       b.Title,
		Description: b.Description,
		Image:       b.Image,
	}
}
