package api

import (
	"net/http"
	"strconv"

	"casita/internal/data"
	"casita/internal/db"

	"github.com/gofiber/fiber/v2"
)

type CreditHandler struct {
	Store db.CreditStore
}

func NewCreditHandler(s *db.Store) *CreditHandler {
	return &CreditHandler{
		Store: s.CreditStore,
	}
}

func (ch *CreditHandler) HandlerGetCredits(c *fiber.Ctx) {
	credits, err := ch.Store.GetAllCredits(c.Context(), nil)
	if err != nil {
		notFoundError(c)
	}

	err = writeJSON(c, http.StatusOK, "got you", credits, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ch *CreditHandler) HandlerGetCredit(c *fiber.Ctx) {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	credit, err := ch.Store.GetCreditByID(c.Context(), nil, id)
	if err != nil {
		notFoundError(c)
	}

	err = writeJSON(c, http.StatusOK, "got you", credit, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ch *CreditHandler) HandlerPostCredit(c *fiber.Ctx) {
	var params *models.CreateCredit
	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
	}

	cred, err := models.NewCreditFromParams(params)
	if err != nil {
		badRequestError(c)
	}

	credResp, err := ch.Store.InsertCredit(c.Context(), nil, cred)
	if err != nil {
		internalServerError(c)
	}

	err = writeJSON(c, http.StatusOK, "got you", credResp, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ch *CreditHandler) HandlerUpdateCredit(c *fiber.Ctx) {
	var params models.UpdateCredit

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
	}

	credResp, err := ch.Store.UpdateCredit(c.Context(), nil, id, &params)
	if err != nil {
		internalServerError(c)
	}

	// TODO: Standardize messages
	err = writeJSON(c, http.StatusOK, "updated successfully", credResp, "")
	if err != nil {
		internalServerError(c)
	}
}

func (ch *CreditHandler) HandlerDeleteCredit(c *fiber.Ctx) {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
	}

	if err := ch.Store.DeleteCreditByID(c.Context(), nil, id); err != nil {
		internalServerError(c)
	}

	err = writeJSON(c, http.StatusOK, "resource deleted", nil, "")
	if err != nil {
		internalServerError(c)
	}
}
