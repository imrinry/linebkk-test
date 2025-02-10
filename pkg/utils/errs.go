package utils

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

var (
	ErrRedisNil = errors.New("redis: nil")
)

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
