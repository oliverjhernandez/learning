package api

import (
	"finance/models"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return ErrUnauthorized()
	}

	if !user.IsAdmin {
		return ErrUnauthorized()
	}

	return c.Next()
}
