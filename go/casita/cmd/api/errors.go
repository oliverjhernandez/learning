package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ApiError struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e ApiError) Error() string {
	return "helloooooo"
}

// func (e ApiError) log(r *http.Request, err error, logger jsonlog.Logger) {
// 	logger.PrintError(e, map[string]string{
// 		"request_method": r.Method,
// 		"request_url":    r.URL.String(),
// 	})
// }

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(*fiber.Error); ok {
		NewApiError(c, apiError.Code, apiError.Message, map[string]string{})
		return apiError
	}
	NewApiError(c, fiber.StatusInternalServerError, err.Error(), map[string]string{})
	return &fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: "the server encountered a problem and could not process your request",
	}
}

func NewApiError(c *fiber.Ctx, status int, message string, errs map[string]string) error {
	err := ApiError{
		Status:  strconv.Itoa(status),
		Message: message,
		Errors:  errs,
	}

	return c.Status(status).JSON(err)
}

func internalServerError(c *fiber.Ctx) error {
	message := "the server encountered a problem and could not process your request"
	return NewApiError(c, fiber.StatusInternalServerError, message, map[string]string{})
}

func notFoundError(c *fiber.Ctx) error {
	message := "the resource you requested could not be found"
	return NewApiError(c, fiber.StatusNotFound, message, map[string]string{})
}

func methodNotAllowedError(c *fiber.Ctx) error {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Method())
	return NewApiError(c, fiber.StatusMethodNotAllowed, message, map[string]string{})
}

func badRequestError(c *fiber.Ctx) error {
	message := "the server was unable to process the request"
	return NewApiError(c, fiber.StatusBadRequest, message, map[string]string{})
}

func editConflictError(c *fiber.Ctx) error {
	message := "unable to update the record due to an edit conflict, please try again"
	return NewApiError(c, fiber.StatusConflict, message, map[string]string{})
}

func unauthorizedError(c *fiber.Ctx) error {
	message := "unauthorized access"
	return NewApiError(c, fiber.StatusUnauthorized, message, map[string]string{})
}

func invalidCredentials(c *fiber.Ctx) error {
	message := "invalid credentials"
	return NewApiError(c, fiber.StatusUnauthorized, message, map[string]string{})
}

func failedValidationResponse(c *fiber.Ctx, errs map[string]string) error {
	message := "validation failed"
	return NewApiError(c, fiber.StatusUnprocessableEntity, message, errs)
}
