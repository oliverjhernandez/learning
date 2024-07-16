package api

import (
	"net/http"
	"strconv"

	"casita/db"
	"casita/models"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	Store db.AccountStore
}

func NewAccountHandler(s *db.Store) *AccountHandler {
	return &AccountHandler{
		Store: s.AccountStore,
	}
}

func (ah *AccountHandler) HandlerPostAccount(c *fiber.Ctx) error {
	var params *models.CreateAccount
	if err := c.BodyParser(&params); err != nil {
		return NewError(http.StatusBadRequest, INVALID_PARAMETERS)
	}

	acc, err := models.NewAccountFromParams(params)
	if err != nil {
		return err
	}

	_, err = ah.Store.InsertAccount(c.Context(), nil, acc)
	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"msg": "updated"})
}

func (ah *AccountHandler) HandlerGetAccounts(c *fiber.Ctx) error {
	accs, err := ah.Store.GetAllAccounts(c.Context(), nil)
	if err != nil {
		return NewError(http.StatusNotFound, NOT_FOUND)
	}
	return c.JSON(accs)
}

func (ah *AccountHandler) HandlerGetAccount(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	acc, err := ah.Store.GetAccountByID(c.Context(), nil, id)
	if err != nil {
		return NewError(http.StatusNotFound, NOT_FOUND)
	}

	return c.JSON(acc)
}

func (ah *AccountHandler) HandlerUpdateAccount(c *fiber.Ctx) error {
	var params models.UpdateAccount

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&params); err != nil {
		return NewError(http.StatusBadRequest, INVALID_REQUEST)
	}

	if err = ah.Store.UpdateAccount(c.Context(), nil, id, &params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "updated"})
}

func (ah *AccountHandler) HandlerDeleteAccount(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := ah.Store.DeleteAccountByID(c.Context(), nil, id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "deleted"})
}
