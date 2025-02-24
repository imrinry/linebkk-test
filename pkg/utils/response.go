package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AppResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type AppPaginationResponse struct {
	Message    string      `json:"message"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Total      int         `json:"total,omitempty"`
	Page       int         `json:"page,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	NextPage   int         `json:"next_page,omitempty"`
	PrevPage   int         `json:"prev_page,omitempty"`
}

// HandleError handles the error and returns the appropriate response
func HandleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case AppError:
		return c.Status(e.Code).JSON(e)
	case error:
		return c.Status(http.StatusInternalServerError).JSON("Internal Server Error")
	default:
		return c.Status(http.StatusInternalServerError).JSON("Internal Server Error")
	}
}

func HandleResponse(c *fiber.Ctx, d interface{}) error {
	switch d := d.(type) {
	case AppResponse:
		return c.Status(d.Code).JSON(d)
	case AppPaginationResponse:
		return c.Status(d.Code).JSON(d)
	default:
		return c.Status(http.StatusInternalServerError).JSON("Internal Server Error")
	}
}
