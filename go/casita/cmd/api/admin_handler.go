package api

import (
	"net/http"

	"casita/internal/data"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return NewError(http.StatusUnauthorized, UNAUTHORIZED)
	}

	if !user.IsAdmin {
		return NewError(http.StatusUnauthorized, UNAUTHORIZED)
	}

	return c.Next()
}
