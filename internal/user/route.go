package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, userHandler UserHandler) {
	userRoutes := app.Group("/users")
	userRoutes.Get("/", userHandler.GetUsers)
}
