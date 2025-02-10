package transactions

import (
	"line-bk-api/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler interface {
	GetTransactionByUserID(ctx *fiber.Ctx) error
}

type transactionHandler struct {
	transactionService TransactionService
}

func NewTransactionHandler(transactionService TransactionService) TransactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

// @Summary Get transactions by user ID
// @Description Get transactions by user ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Security ApiKeyAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} utils.AppPaginationResponse{data=[]TransactionResponse,total=int,page=int,limit=int}
// @Failure 404 {object} utils.AppError{message=string,code=int}
// @Failure 500 {object} utils.AppError{message=string,code=int}
// @Router /api/v1/transactions [get]
func (h *transactionHandler) GetTransactionByUserID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	transactions, total, err := h.transactionService.GetTransactionByUserID(c.Context(), userID, page, limit)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleResponse(c, utils.AppPaginationResponse{
		Message:    "success",
		Code:       http.StatusOK,
		Data:       transactions,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: utils.GetTotalPages(total, limit),
		NextPage:   utils.GetNextPage(page, utils.GetTotalPages(total, limit)),
		PrevPage:   utils.GetPreviousPage(page),
	})
}
