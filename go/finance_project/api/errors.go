package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"statusCode"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.StatusCode).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.StatusCode).JSON(apiError)
}

func (e Error) Error() string {
	return e.Message
}

func NewError(statusCode int, message string) Error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func ErrUnauthorized() Error {
	return Error{
		StatusCode: http.StatusUnauthorized,
		Message:    "unauthorized",
	}
}

func ErrNotFound() Error {
	return Error{
		StatusCode: http.StatusNotFound,
		Message:    "not found",
	}
}

func ErrInvalidParams() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid parameters",
	}
}

func ErrInvalidReqBody() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid request body",
	}
}

func ErrInvalidID() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid id",
	}
}
