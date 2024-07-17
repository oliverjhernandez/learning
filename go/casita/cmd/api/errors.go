package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(*fiber.Error); ok {
		NewApiError(c, fiber.StatusInternalServerError, apiError.Message)
		return apiError
	}
	NewApiError(c, fiber.StatusInternalServerError, err.Error())
	return &fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: "the server encountered a problem and could not process your request",
	}
}

func NewApiError(c *fiber.Ctx, status int, error string) {
	err := Envelope{
		Status: strconv.Itoa(status),
		Error:  error,
	}

	c.Response().Header.Add("Content-Type", "application/json")
	c.Status(status).JSON(err)
}

func internalServerError(c *fiber.Ctx) *fiber.Error {
	message := "the server encountered a problem and could not process your request"
	NewApiError(c, fiber.StatusInternalServerError, message)
	return &fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	}
}

func notFoundError(c *fiber.Ctx) *fiber.Error {
	message := "the resource you requested could not be found"
	NewApiError(c, fiber.StatusNotFound, message)
	return &fiber.Error{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}

func methodNotAllowedError(c *fiber.Ctx) *fiber.Error {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Method())
	NewApiError(c, fiber.StatusMethodNotAllowed, message)
	return &fiber.Error{
		Code:    fiber.StatusMethodNotAllowed,
		Message: message,
	}
}

func badRequestError(c *fiber.Ctx) *fiber.Error {
	message := "the server was unable to process the request"
	NewApiError(c, fiber.StatusBadRequest, message)
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}

func editConflictError(c *fiber.Ctx) *fiber.Error {
	message := "unable to update the record due to an edit conflict, please try again"
	NewApiError(c, fiber.StatusConflict, message)
	return &fiber.Error{
		Code:    fiber.StatusConflict,
		Message: message,
	}
}

func unauthorizedError(c *fiber.Ctx) *fiber.Error {
	message := "unauthorized access"
	NewApiError(c, fiber.StatusUnauthorized, message)
	return &fiber.Error{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	}
}

func invalidCredentials(c *fiber.Ctx) *fiber.Error {
	message := "invalid credentials"
	NewApiError(c, fiber.StatusUnauthorized, message)
	return &fiber.Error{
		Code: fiber.StatusUnauthorized,
	}
}
