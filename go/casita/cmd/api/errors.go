package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewError(c *fiber.Ctx, status int, error string) {
	err := Envelope{
		Status: strconv.Itoa(status),
		Error:  error,
	}

	c.Response().Header.Add("Content-Type", "application/json")

	c.Status(status).JSON(err)
}

func internalServerError(c *fiber.Ctx) {
	message := "the server encountered a problem and could not process your request"
	NewError(c, http.StatusInternalServerError, message)
}

func notFoundError(c *fiber.Ctx) {
	message := "the resource you requested could not be found"
	NewError(c, http.StatusNotFound, message)
}

func methodNotAllowedError(c *fiber.Ctx) {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Method())
	NewError(c, http.StatusMethodNotAllowed, message)
}

func badRequestError(c *fiber.Ctx) {
	message := "the server was unable to process the request"
	NewError(c, http.StatusBadRequest, message)
}

func editConflictError(c *fiber.Ctx) {
	message := "unable to update the record due to an edit conflict, please try again"
	NewError(c, http.StatusConflict, message)
}

func unauthorizedError(c *fiber.Ctx) {
	message := "unauthorized access"
	NewError(c, http.StatusUnauthorized, message)
}

func invalidCredentials(c *fiber.Ctx) {
	message := "invalid credentials"
	NewError(c, http.StatusUnauthorized, message)
}
