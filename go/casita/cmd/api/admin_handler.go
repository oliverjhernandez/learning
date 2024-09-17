package api

import (
	"errors"
	"net/http"

	"casita/internal/models"
)

func AdminAuth(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	user, ok := c.Value("user").(*models.User)
	if !ok {
		unauthorizedError(errors.New("unauthorizedError"))
	}

	if !*user.IsAdmin {
		unauthorizedError(errors.New("unauthorizedError"))
	}

	// c.Next()
}
