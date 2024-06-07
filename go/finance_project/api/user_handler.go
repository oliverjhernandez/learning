package api

import (
	"finance/db"
	"finance/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStore db.UserStore
}

func (uh *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	users, err := uh.userStore.GetUsers(c.Context())
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(users)
}

func (uh *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uh.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(user)
}

func (uh *UserHandler) HandlerPostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	res, err := uh.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (uh *UserHandler) HandlerDeleteUser(c *fiber.Ctx) error {
	if err := uh.userStore.DeleteUser(c.Context(), c.Params("id")); err != nil {
		return err
	}

	return nil
}

func (uh *UserHandler) HandlerUpdateUser(c *fiber.Ctx) error {
	var params types.UpdateUserParams
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	oid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return ErrInvalidID()
	}
	filter := bson.M{"_id": oid}

	if err := uh.userStore.UpdateUser(c.Context(), filter, &params); err != nil {
		return err
	}
	return nil
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}
