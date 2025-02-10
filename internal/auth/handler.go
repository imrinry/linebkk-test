package auth

import (
	"line-bk-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	LoginWithPinCode(c *fiber.Ctx) error
	LoginWithPassword(c *fiber.Ctx) error
}

type authHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

// @Summary Login with pin code
// @Description Login with pin code
// @Accept json
// @Tags Authentication
// @Produce json
// @Security ApiKeyAuth
// @Param loginWithPinCodeRequest body LoginWithPinCodeRequest true "Login with pin code request"
// @Success 200 {object} utils.AppResponse{data=LoginResponse}
// @Failure 400 {object} utils.AppError{message=string}
// @Failure 401 {object} utils.AppError{message=string}
// @Failure 500 {object} utils.AppError{message=string}
// @Router /api/v1/auth/login/pin [post]
func (h *authHandler) LoginWithPinCode(c *fiber.Ctx) error {
	var request LoginWithPinCodeRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.HandleError(c, err)
	}

	data, err := h.authService.LoginWithPinCode(request.UserID, request.PinCode)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.AppResponse{
		Data: data,
	})
}

// @Summary Login with password
// @Description Login with password
// @Accept json
// @Tags Authentication
// @Produce json
// @Security ApiKeyAuth
// @Param loginWithPasswordRequest body LoginWithPasswordRequest true "Login with password request"
// @Success 200 {object} utils.AppResponse{data=LoginResponse}
// @Failure 400 {object} utils.AppError{message=string}
// @Failure 401 {object} utils.AppError{message=string}
// @Failure 500 {object} utils.AppError{message=string}
// @Router /api/v1/auth/login/password [post]
func (h *authHandler) LoginWithPassword(c *fiber.Ctx) error {
	var request LoginWithPasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.HandleError(c, err)
	}

	data, err := h.authService.LoginWithPassword(request.UserID, request.Password)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.AppResponse{
		Data:    data,
		Code:    200,
		Message: "Login successful",
	})
}
