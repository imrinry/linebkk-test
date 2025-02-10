package debit_cards

import (
	"line-bk-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterDebitCardRoutes(app *fiber.App, debitCardHandler DebitCardHandler) {
	api := app.Group("/api/v1")

	api.Use(middleware.AuthMiddleware())
	api.Get("/debit-cards", debitCardHandler.GetDebitCards)
}
