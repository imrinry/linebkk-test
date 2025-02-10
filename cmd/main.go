package main

import (
	"line-bk-api/config"
	"line-bk-api/internal/account"
	"line-bk-api/internal/auth"
	"line-bk-api/internal/banner"
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
// @description Enter the token with the prefix **"Bearer "**, e.g., "Bearer {your_token}"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
// @description The API key for the LINE BK API.
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

	accountRepo := account.NewAccountRepository(config.GetDBInstance())
	accountService := account.NewAccountService(accountRepo)
	accountHandler := account.NewHandler(accountService)

	bannerRepo := banner.NewBannerRepository(config.GetDBInstance())
	bannerService := banner.NewBannerService(bannerRepo)
	bannerHandler := banner.NewBannerHandler(bannerService)

	routes.SetupRoutes(app, userHandler, authHandler, accountHandler, bannerHandler)

	log.Fatal(app.Listen(":8080"))
}
