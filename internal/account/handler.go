package account

import (
	"line-bk-api/config"
	"line-bk-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	accountService AccountService
}

func NewHandler(accountService AccountService) Handler {
	return Handler{accountService: accountService}
}

// @Summary Get account my account
// @Description Get account my account
// @Accept json
// @Produce json
// @Tags Account
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} utils.AppPaginationResponse{data=[]AccountResponse,total=int,page=int,limit=int,total_pages=int,next_page=int,prev_page=int}
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /api/v1/accounts/me [get]
func (h *Handler) GetMyAccount(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page := c.QueryInt("page", config.DefaultPage)
	limit := c.QueryInt("limit", config.DefaultLimit)

	accounts, total, err := h.accountService.GetAccountByUserID(userID, page, limit)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleResponse(c, utils.AppPaginationResponse{
		Message:    "success",
		Code:       fiber.StatusOK,
		Data:       accounts,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: utils.GetTotalPages(total, limit),
		NextPage:   utils.GetNextPage(page, utils.GetTotalPages(total, limit)),
		PrevPage:   utils.GetPreviousPage(page),
	})
}
