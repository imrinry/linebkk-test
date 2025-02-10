package account

import (
	"line-bk-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAccountRoutes(app *fiber.App, accountHandler Handler) {
	api := app.Group("/api/v1")

	api.Use(middleware.AuthMiddleware())
	api.Get("/accounts/me", accountHandler.GetMyAccount)
}
