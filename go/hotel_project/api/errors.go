package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Msg  string `json:"message"`
	Code int    `json:"statusCode"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

func (e Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

func ErrUnauthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Msg:  "unauthorized",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Msg:  "invalid request",
	}
}

func ErrNotFound() Error {
	return Error{
		Code: http.StatusNotFound,
		Msg:  "not found",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Msg:  "invalid id given",
	}
}
