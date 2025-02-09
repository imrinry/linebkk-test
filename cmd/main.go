package main

import (
	"line-bk-api/config"
	"line-bk-api/internal/user"
	"line-bk-api/pkg/logs"
	"line-bk-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDB()

	// Connect to Redis
	// config.ConnectRedis()

	app := fiber.New()

	userRepo := user.NewUserRepository(config.GetDBInstance())
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// Setup routes
	routes.SetupRoutes(app, userHandler)

	app.Listen(":8080")
	logs.Info("Server is running on port 8080")
}
