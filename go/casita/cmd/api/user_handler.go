package api

import (
	"net/http"
	"strconv"

	"casita/internal/data"
	"casita/internal/db"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Store db.UserStore
}

func NewUserHandler(s *db.Store) *UserHandler {
	return &UserHandler{
		Store: s.UserStore,
	}
}

func (uh *UserHandler) HandlerPostUser(c *fiber.Ctx) {
	var params models.CreateUser
	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
	}

	user, err := models.NewUserFromParams(&params)
	if err != nil {
		badRequestError(c)
	}

	userResp, err := uh.Store.InsertUser(c.Context(), nil, user)
	if err != nil {
		notFoundError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", userResp, "")
	if err != nil {
		internalServerError(c)
	}
}

func (uh *UserHandler) HandlerGetUsers(c *fiber.Ctx) {
	users, err := uh.Store.GetAllUsers(c.Context(), nil)
	if err != nil {
		notFoundError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", users, "")
	if err != nil {
		internalServerError(c)
	}
}

func (uh *UserHandler) HandlerGetUser(c *fiber.Ctx) {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	user, err := uh.Store.GetUserByID(c.Context(), nil, id)
	if err != nil {
		badRequestError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", user, "")
	if err != nil {
		internalServerError(c)
	}
}

func (uh *UserHandler) HandlerUpdateUser(c *fiber.Ctx) {
	var params models.UpdateUser

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
	}

	userResp, err := uh.Store.UpdateUser(c.Context(), nil, id, &params)
	if err != nil {
		internalServerError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", userResp, "")
	if err != nil {
		internalServerError(c)
	}
}

func (uh *UserHandler) HandlerDeleteUser(c *fiber.Ctx) {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	if err := uh.Store.DeleteUserByID(c.Context(), nil, id); err != nil {
		internalServerError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource deleted successfully", nil, "")
	if err != nil {
		internalServerError(c)
	}
}
