package debit_cards

import (
	"line-bk-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type DebitCardHandler interface {
	GetDebitCards(c *fiber.Ctx) error
}

type debitCardHandler struct {
	debitCardService DebitCardService
}

func NewDebitCardHandler(debitCardService DebitCardService) DebitCardHandler {
	return &debitCardHandler{debitCardService: debitCardService}
}

// @Summary Get debit cards
// @Description Get debit cards
// @Accept json
// @Produce json
// @Tags DebitCard
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} utils.AppPaginationResponse{data=[]DebitCardResponse,total=int,page=int,limit=int,total_pages=int,next_page=int,prev_page=int}
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /api/v1/debit-cards [get]
func (h *debitCardHandler) GetDebitCards(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	debitCards, total, err := h.debitCardService.GetDebitCards(c.Context(), userID, page, limit)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleResponse(c, utils.AppPaginationResponse{
		Message:    "success",
		Code:       fiber.StatusOK,
		Data:       debitCards,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: utils.GetTotalPages(total, limit),
		NextPage:   utils.GetNextPage(page, utils.GetTotalPages(total, limit)),
		PrevPage:   utils.GetPreviousPage(page),
	})
}
