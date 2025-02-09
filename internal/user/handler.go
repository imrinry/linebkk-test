package user

import (
	"fmt"
	"line-bk-api/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	GetMyProfile(c *fiber.Ctx) error
}

type handler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &handler{userService: userService}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a paginated list of users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} utils.AppPaginationResponse{data=[]UserResponseDTO,total=int,page=int,limit=int}
// @Failure 500 {object} utils.AppError{message=string,code=int}
// @Router /api/v1/users [get]
func (h *handler) GetUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	limit := c.QueryInt("limit")

	limit = utils.GetLimit(limit)

	users, total, err := h.userService.GetUsers(page, limit)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleResponse(c, utils.AppPaginationResponse{
		Message:    "success",
		Code:       http.StatusOK,
		Data:       users,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: utils.GetTotalPages(total, limit),
		NextPage:   utils.GetNextPage(page, utils.GetTotalPages(total, limit)),
		PrevPage:   utils.GetPreviousPage(page),
	})
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} utils.AppResponse{data=UserResponseDTO}
// @Failure 404 {object} utils.AppError{message=string,code=int}
// @Failure 500 {object} utils.AppError{message=string,code=int}
// @Router /api/v1/users/{id} [get]
func (h *handler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	fmt.Println("userID", userID)

	user, err := h.userService.GetUserByID("userID")
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleResponse(c, utils.AppResponse{
		Message: "success",
		Code:    http.StatusOK,
		Data:    user,
	})
}

// GetMyProfile godoc
// @Summary Get my profile
// @Description Get my profile
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.AppResponse{data=UserResponseDTO}
// @Failure 404 {object} utils.AppError{message=string,code=int}
// @Failure 500 {object} utils.AppError{message=string,code=int}
// @Router /api/v1/users/me [get]
func (h *handler) GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleResponse(c, utils.AppResponse{
		Message: "success",
		Code:    http.StatusOK,
		Data:    user,
	})
}
