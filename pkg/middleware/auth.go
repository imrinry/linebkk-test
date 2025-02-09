package middleware

import (
	"fmt"
	"line-bk-api/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	fmt.Println("\n=== Auth Middleware Starting ===")
	fmt.Println("Method:", c.Method())
	fmt.Println("Path:", c.Path())
	fmt.Println("Full URL:", c.BaseURL()+c.OriginalURL())

	token := c.Get("Authorization")
	fmt.Println("Received token:", token)

	if token == "" {
		fmt.Println("No token provided - returning 401")
		return utils.AppError{
			Message: "No authorization token provided",
			Code:    401,
		}
	}

	// Check and process Bearer token
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		fmt.Println("Processed Bearer token:", token)
	}

	// Validate token regardless of Bearer prefix
	userID, err := utils.ValidateAccessToken(token)
	if err != nil {
		fmt.Println("Token validation failed:", err)
		return utils.AppError{
			Message: "Unauthorized",
			Code:    401,
		}
	}

	fmt.Println("Token validated successfully for user:", userID)
	c.Locals("user_id", userID)
	fmt.Println("=== Auth Middleware Completed ===\n")
	return c.Next()
}
