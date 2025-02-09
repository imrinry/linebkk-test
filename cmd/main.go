package main

import (
	"line-bk-api/config"
	"line-bk-api/internal/auth"
	"line-bk-api/internal/user"
	"line-bk-api/routes"
	"log"
	"time"

	_ "line-bk-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title LINE BK API
// @version 1.0
// @description This is a sample server for LINE BK API.
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @scheme bearer
func main() {
	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		log.Printf(
			"[%s] %s %s - %d (%v)",
			c.Method(),
			c.Path(),
			c.IP(),
			c.Response().StatusCode(),
			time.Since(start),
		)

		return err
	})

	// Swagger route at root level
	app.Get("/swagger/*", swagger.HandlerDefault)

	userRepo := user.NewUserRepository(config.GetDBInstance())
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	authRepo := auth.NewAuthRepository(config.GetDBInstance())
	authService := auth.NewAuthService(authRepo, userRepo)
	authHandler := auth.NewAuthHandler(authService)


	routes.SetupRoutes(app, userHandler, authHandler)

	log.Fatal(app.Listen(":8080"))
}
