package routes

import (
	"line-bk-api/internal/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler user.UserHandler) {
	user.RegisterRoutes(app, userHandler)
}
