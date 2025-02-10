package transactions

import (
	"line-bk-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterTransactionRoutes(router *fiber.App, transactionHandler TransactionHandler) {
	v1 := router.Group("/api/v1")
	transactionRouter := v1.Group("/transactions")

	transactionRouter.Use(middleware.AuthMiddleware())
	transactionRouter.Get("/", transactionHandler.GetTransactionByUserID)
}
