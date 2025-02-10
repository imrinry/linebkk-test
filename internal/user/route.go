package user

import (
	"line-bk-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userHandler UserHandler) {
	v1 := app.Group("/api/v1")
	userRoutes := v1.Group("/users")

	// todo : for admin dashboard
	// userRoutes.Get("/", userHandler.GetUsers)
	// userRoutes.Get("/:id", userHandler.GetUserByID)

	// Protected routes
	userRoutes.Use(middleware.AuthMiddleware())
	userRoutes.Get("/profile/me", userHandler.GetMyProfile)
	userRoutes.Get("/greetings", userHandler.GetUserGreeting)
}
