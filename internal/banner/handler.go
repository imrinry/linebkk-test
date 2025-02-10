package banner

import (
	"line-bk-api/config"
	"line-bk-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type BannerHandler interface {
	GetBannerByUserID(c *fiber.Ctx) error
}

type bannerHandler struct {
	bannerService BannerService
}

func NewBannerHandler(bannerService BannerService) BannerHandler {
	return &bannerHandler{bannerService: bannerService}
}

// @Summary Get banner by user id
// @Description Get banner by user id
// @Accept json
// @Produce json
// @Tags Banner
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} utils.AppPaginationResponse{data=[]BannerResponse,total=int,page=int,limit=int,total_pages=int,next_page=int,prev_page=int}
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /api/v1/banners [get]
func (h *bannerHandler) GetBannerByUserID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page := c.QueryInt("page", config.DefaultPage)
	limit := c.QueryInt("limit", config.DefaultLimit)

	banners, total, err := h.bannerService.GetBannerByUserID(c.Context(), userID, page, limit)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleResponse(c, utils.AppPaginationResponse{
		Message:    "success",
		Code:       fiber.StatusOK,
		Data:       banners,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: utils.GetTotalPages(total, limit),
		NextPage:   utils.GetNextPage(page, utils.GetTotalPages(total, limit)),
		PrevPage:   utils.GetPreviousPage(page),
	})
}
