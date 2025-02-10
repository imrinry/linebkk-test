package auth

import "github.com/gofiber/fiber/v2"

func RegisterAuthRoutes(app *fiber.App, authHandler AuthHandler) {
	v1 := app.Group("/api/v1")
	authRoutes := v1.Group("/auth")
	authRoutes.Post("/login/pin", authHandler.LoginWithPinCode)
	authRoutes.Post("/login/password", authHandler.LoginWithPassword)
}
