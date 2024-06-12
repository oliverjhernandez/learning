package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorMessage int

const (
	_ = iota
	UNAUTHORIZED
	NOT_FOUND
	INVALID_PARAMETERS
	INVALID_ID
	INVALID_REQUEST_BODY
)

func (e ErrorMessage) String() string {
	switch e {
	case UNAUTHORIZED:
		return "UNAUTHORIZED"
	case NOT_FOUND:
		return "NOT_FOUND"
	case INVALID_PARAMETERS:
		return "INVALID_PARAMETERS"
	case INVALID_ID:
		return "INVALID_ID"
	case INVALID_REQUEST_BODY:
		return "INVALID_REQUEST_BODY"
	default:
		return "Unknown"
	}
}

type Error struct {
	Message    string       `json:"msg"` // TODO: Maybe add the original error message
	ErrMessage ErrorMessage `json:"errMsg"`
	StatusCode int          `json:"statusCode"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.StatusCode).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.StatusCode).JSON(apiError)
}

func (e Error) Error() string {
	return e.ErrMessage.String()
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
		ErrMessage: UNAUTHORIZED,
	}
}

func ErrNotFound() Error {
	return Error{
		StatusCode: http.StatusNotFound,
		ErrMessage: NOT_FOUND,
	}
}

func ErrInvalidParams() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		ErrMessage: INVALID_PARAMETERS,
	}
}

func ErrInvalidReqBody() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		ErrMessage: INVALID_REQUEST_BODY,
	}
}

func ErrInvalidID() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		ErrMessage: INVALID_ID,
	}
}
