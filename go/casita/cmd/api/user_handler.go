package api

import (
	"net/http"
	"strconv"

	"casita/internal/db"
	"casita/internal/models"

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
	if err := readJSON(c, &params); err != nil {
		badRequestError(c)
		return err
	}

	user, err := models.NewUserFromParams(&params)
	if err != nil {
		badRequestError(c)
		return err
	}

	userResp, err := uh.Store.InsertUser(c.Context(), nil, user)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", userResp, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (uh *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	users, err := uh.Store.GetAllUsers(c.Context(), nil)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", users, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (uh *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	user, err := uh.Store.GetUserByID(c.Context(), nil, id)
	if err != nil {
		badRequestError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", user, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (uh *UserHandler) HandlerUpdateUser(c *fiber.Ctx) error {
	var params models.UpdateUser

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	if err := readJSON(c, &params); err != nil {
		badRequestError(c)
		return err
	}

	userResp, err := uh.Store.UpdateUser(c.Context(), nil, id, &params)
	if err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", userResp, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (uh *UserHandler) HandlerDeleteUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	if err := uh.Store.DeleteUserByID(c.Context(), nil, id); err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource deleted successfully", nil, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}
