package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
}

type handler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &handler{userService: userService}
}

func (h *handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers(c.QueryInt("page"), c.QueryInt("limit"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.JSON(users)
}
