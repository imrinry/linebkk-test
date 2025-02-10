package routes

import (
	"line-bk-api/internal/account"
	"line-bk-api/internal/auth"
	"line-bk-api/internal/banner"
	"line-bk-api/internal/debit_cards"
	"line-bk-api/internal/transactions"
	"line-bk-api/internal/user"
	"line-bk-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler user.UserHandler, authHandler auth.AuthHandler, accountHandler account.Handler, bannerHandler banner.BannerHandler, transactionHandler transactions.TransactionHandler, debitCardHandler debit_cards.DebitCardHandler) {
	app.Use(middleware.CheckApiKey())
	user.RegisterUserRoutes(app, userHandler)
	auth.RegisterAuthRoutes(app, authHandler)
	account.RegisterAccountRoutes(app, accountHandler)
	banner.RegisterBannerRoutes(app, bannerHandler)
	transactions.RegisterTransactionRoutes(app, transactionHandler)
	debit_cards.RegisterDebitCardRoutes(app, debitCardHandler)
}
