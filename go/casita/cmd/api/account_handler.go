package api

import (
	"net/http"
	"strconv"

	"casita/internal/data"
	"casita/internal/db"

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

func (ah *AccountHandler) HandlerPostAccount(c *fiber.Ctx) {
	var params *models.CreateAccount
	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
	}

	acc, err := models.NewAccountFromParams(params)
	if err != nil {
		badRequestError(c)
	}

	acc, err = ah.Store.InsertAccount(c.Context(), nil, acc)
	if err != nil {
		internalServerError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", acc, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ah *AccountHandler) HandlerGetAccounts(c *fiber.Ctx) {
	accs, err := ah.Store.GetAllAccounts(c.Context(), nil)
	if err != nil {
		notFoundError(c)
	}

	err = writeJSON(c, http.StatusOK, "got you", accs, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ah *AccountHandler) HandlerGetAccount(c *fiber.Ctx) {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	acc, err := ah.Store.GetAccountByID(c.Context(), nil, id)
	if err != nil {
		notFoundError(c)
	}

	err = writeJSON(c, http.StatusOK, "got you", acc, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ah *AccountHandler) HandlerUpdateAccount(c *fiber.Ctx) {
	var params models.UpdateAccount

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
	}

	acc, err := ah.Store.UpdateAccount(c.Context(), nil, id, &params)
	if err != nil {
		editConflictError(c)
	}

	err = writeJSON(c, http.StatusOK, "updated successfully", acc, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ah *AccountHandler) HandlerDeleteAccount(c *fiber.Ctx) {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	if err := ah.Store.DeleteAccountByID(c.Context(), nil, id); err != nil {
		internalServerError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource deleted", nil, "")
	if err != nil {
		internalServerError(c)
	}
}
