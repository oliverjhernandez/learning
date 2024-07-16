package api

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorCode int

const (
	_ = iota
	UNAUTHORIZED
	NOT_FOUND
	INVALID_PARAMETERS
	INVALID_ID
	INVALID_REQUEST
	INTERNAL_ERROR
	TOKEN_EXPIRED
	INVALID_TOKEN
)

func (e ErrorCode) String() string {
	switch e {
	case UNAUTHORIZED:
		return "UNAUTHORIZED"
	case NOT_FOUND:
		return "NOT_FOUND"
	case INVALID_PARAMETERS:
		return "INVALID_PARAMETERS"
	case INVALID_ID:
		return "INVALID_ID"
	case INVALID_REQUEST:
		return "INVALID_REQUEST"
	case INTERNAL_ERROR:
		return "INTERNAL_ERROR"
	case TOKEN_EXPIRED:
		return "TOKEN_EXPIRED"
	case INVALID_TOKEN:
		return "INVALID_TOKEN"
	default:
		return "UNKNOWN"
	}
}

func (e ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

type Error struct {
	ErrMessage ErrorCode `json:"error"`
	StatusCode int       `json:"status_code"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.StatusCode).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, INTERNAL_ERROR)
	return c.Status(apiError.StatusCode).JSON(apiError)
}

func (e Error) Error() string {
	return e.ErrMessage.String()
}

func NewError(statusCode int, error ErrorCode) Error {
	return Error{
		StatusCode: statusCode,
		ErrMessage: error,
	}
}
