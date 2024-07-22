package api

import (
	"net/http"
	"strconv"

	"casita/internal/db"
	"casita/internal/models"

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

var params *models.CreateAccount

func (ah *AccountHandler) HandlerPostAccount(c *fiber.Ctx) error {
	var params models.CreateAccount
	if err := readJSON(c, &params); err != nil {
		badRequestError(c)
		return err
	}

	acc, err := models.NewAccountFromParams(&params)
	if err != nil {
		badRequestError(c)
		return err
	}

	acc, err = ah.Store.InsertAccount(c.Context(), nil, acc)
	if err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", acc, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ah *AccountHandler) HandlerGetAccounts(c *fiber.Ctx) error {
	accs, err := ah.Store.GetAllAccounts(c.Context(), nil)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", accs, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ah *AccountHandler) HandlerGetAccount(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	acc, err := ah.Store.GetAccountByID(c.Context(), nil, id)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", acc, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ah *AccountHandler) HandlerUpdateAccount(c *fiber.Ctx) error {
	var params models.UpdateAccount

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

	acc, err := ah.Store.UpdateAccount(c.Context(), nil, id, &params)
	if err != nil {
		editConflictError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "updated successfully", acc, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ah *AccountHandler) HandlerDeleteAccount(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	if err := ah.Store.DeleteAccountByID(c.Context(), nil, id); err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource deleted", nil, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}
