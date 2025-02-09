package routes

import (
	"line-bk-api/internal/auth"
	"line-bk-api/internal/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler user.UserHandler, authHandler auth.AuthHandler) {
	user.RegisterRoutes(app, userHandler)
	auth.RegisterAuthRoutes(app, authHandler)
}
