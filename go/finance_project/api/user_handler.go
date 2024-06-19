package api

import (
	"strconv"

	"finance/db"
	"finance/models"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Store *db.PGUserStore
}

func (uh *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	users, err := uh.Store.GetAllUsers()
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(users)
}

func (uh *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	user, err := uh.Store.GetUserByID(id)
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(user)
}

func (uh *UserHandler) HandlerPostUser(c *fiber.Ctx) error {
	var params models.CreateUser
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	user := models.NewUserFromParams(params)

	res, err := uh.Store.InsertUser(user)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (uh *UserHandler) HandlerDeleteUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := uh.Store.DeleteUserByID(id); err != nil {
		return err
	}

	return nil
}

func (uh *UserHandler) HandlerUpdateUser(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	var params models.UpdateUser
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	if err := uh.Store.UpdateUser(id, &params); err != nil {
		return err
	}
	return nil
}

func NewUserHandler(s *db.PGUserStore) *UserHandler {
	return &UserHandler{
		Store: s,
	}
}
