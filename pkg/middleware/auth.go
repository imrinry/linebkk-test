package middleware

import (
	"line-bk-api/config"
	"line-bk-api/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Get("Authorization")

		if token == "" {
			return utils.AppError{
				Message: "No authorization token provided",
				Code:    401,
			}
		}

		// Check and process Bearer token
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
			userID, err := utils.ValidateAccessToken(token)
			if err != nil {
				if err.Error() == "Token is expired" {
					return utils.HandleError(c, utils.AppError{
						Message: "Token is expired",
						Code:    401,
					})
				}
				return utils.HandleError(c, err)
			}
			c.Locals("user_id", userID)
		} else {
			userID, err := utils.ValidateAccessToken(token)
			if err != nil {
				if err.Error() == "Token is expired" {
					return utils.HandleError(c, utils.AppError{
						Message: "Token is expired",
						Code:    401,
					})
				}
				return utils.HandleError(c, err)
			}
			c.Locals("user_id", userID)
		}
		return c.Next()
	}
}

func CheckApiKey() fiber.Handler {
	return func(c *fiber.Ctx) error {

		apiKey := c.Get("X-API-KEY")
		if apiKey != config.X_API_KEY {
			return utils.HandleError(c, utils.AppError{
				Message: "Invalid API key",
				Code:    401,
			})
		}
		return c.Next()
	}
}
