package api

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(*fiber.Error); ok {
		NewApiError(c, apiError.Code, apiError.Message)
		return apiError
	}
	NewApiError(c, fiber.StatusInternalServerError, err.Error())
	return &fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: "the server encountered a problem and could not process your request",
	}
}

func NewApiError(c *fiber.Ctx, status int, message interface{}) {
	err := Envelope{
		Status: strconv.Itoa(status),
		Error:  message,
	}

	c.Response().Header.Add("Content-Type", "application/json")
	c.Status(status).JSON(err)
}

func internalServerError(c *fiber.Ctx) error {
	message := "the server encountered a problem and could not process your request"
	NewApiError(c, fiber.StatusInternalServerError, message)
	return errors.New(message)
}

func notFoundError(c *fiber.Ctx) error {
	message := "the resource you requested could not be found"
	NewApiError(c, fiber.StatusNotFound, message)
	return errors.New(message)
}

func methodNotAllowedError(c *fiber.Ctx) error {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Method())
	NewApiError(c, fiber.StatusMethodNotAllowed, message)
	return errors.New(message)
}

func badRequestError(c *fiber.Ctx) error {
	message := "the server was unable to process the request"
	NewApiError(c, fiber.StatusBadRequest, message)
	return errors.New(message)
}

func editConflictError(c *fiber.Ctx) error {
	message := "unable to update the record due to an edit conflict, please try again"
	NewApiError(c, fiber.StatusConflict, message)
	return errors.New(message)
}

func unauthorizedError(c *fiber.Ctx) error {
	message := "unauthorized access"
	NewApiError(c, fiber.StatusUnauthorized, message)
	return errors.New(message)
}

func invalidCredentials(c *fiber.Ctx) error {
	message := "invalid credentials"
	NewApiError(c, fiber.StatusUnauthorized, message)
	return errors.New(message)
}

func failedValidationResponse(c *fiber.Ctx, message map[string]string) error {
	explanation := "validation failed"
	NewApiError(c, fiber.StatusUnprocessableEntity, message)
	return errors.New(explanation)
}
