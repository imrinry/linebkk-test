package banner

import "github.com/gofiber/fiber/v2"

func RegisterBannerRoutes(app *fiber.App, bannerHandler BannerHandler) {
	v1 := app.Group("/api/v1")
	bannerRoutes := v1.Group("/banners")
	bannerRoutes.Get("/", bannerHandler.GetBannerByUserID)
}
