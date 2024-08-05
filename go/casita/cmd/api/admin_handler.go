package api

import (
	"casita/internal/models"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		unauthorizedError(c)
	}

	if !*user.IsAdmin {
		unauthorizedError(c)
	}

	c.Next()
}
