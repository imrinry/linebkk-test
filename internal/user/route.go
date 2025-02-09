package user

import (
	"line-bk-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, userHandler UserHandler) {
	v1 := app.Group("/api/v1")
	userRoutes := v1.Group("/users")

	// Public routes
	userRoutes.Get("/", userHandler.GetUsers)
	userRoutes.Get("/:id", userHandler.GetUserByID)

	// Protected routes

	userRoutes.Get("/me", userHandler.GetMyProfile, middleware.AuthMiddleware)
}
