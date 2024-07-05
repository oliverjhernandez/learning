package api

import (
	"net/http"
	"strconv"

	"casita/db"
	"casita/models"

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

func (uh *UserHandler) HandlerPostUser(c *fiber.Ctx) error {
	var params models.CreateUser
	if err := c.BodyParser(&params); err != nil {
		return NewError(http.StatusBadRequest, INVALID_PARAMETERS)
	}

	user, err := models.NewUserFromParams(&params)
	if err != nil {
		return NewError(http.StatusBadRequest, INVALID_REQUEST)
	}

	_, err = uh.Store.InsertUser(c.Context(), nil, user)
	if err != nil {
		return NewError(http.StatusNotFound, NOT_FOUND)
	}

	return c.JSON(map[string]string{"msg": "created"})
}

func (uh *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	users, err := uh.Store.GetAllUsers(c.Context(), nil)
	if err != nil {
		return NewError(http.StatusNotFound, NOT_FOUND)
	}
	return c.JSON(users)
}

func (uh *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	user, err := uh.Store.GetUserByID(c.Context(), nil, id)
	if err != nil {
		return NewError(http.StatusNotFound, NOT_FOUND)
	}
	return c.JSON(user)
}

func (uh *UserHandler) HandlerUpdateUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	var params models.UpdateUser
	if err := c.BodyParser(&params); err != nil {
		return NewError(http.StatusBadRequest, INVALID_REQUEST)
	}

	if err := uh.Store.UpdateUser(c.Context(), nil, id, &params); err != nil {
		return err
	}
	return nil
}

func (uh *UserHandler) HandlerDeleteUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := uh.Store.DeleteUserByID(c.Context(), nil, id); err != nil {
		return err
	}

	return nil
}
