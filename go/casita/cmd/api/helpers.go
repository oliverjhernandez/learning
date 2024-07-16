package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Envelope struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func writeJSON(c *fiber.Ctx, status int, message string, data interface{}, error string) error {
	response := Envelope{
		Status:  strconv.Itoa(status),
		Message: message,
		Data:    data,
		Error:   error,
	}

	c.Response().Header.Add("Content-Type", "application/json")

	if err := c.Status(status).JSON(response); err != nil {
		return err
	}

	return nil
}
