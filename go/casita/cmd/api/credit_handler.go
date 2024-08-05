package api

import (
	"net/http"
	"strconv"

	"casita/internal/db"
	"casita/internal/models"
	"casita/internal/validator"

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

func (ch *CreditHandler) HandlerGetCredits(c *fiber.Ctx) error {
	credits, err := ch.Store.GetAllCredits(c.Context(), nil)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", credits, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ch *CreditHandler) HandlerGetCredit(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	credit, err := ch.Store.GetCreditByID(c.Context(), nil, id)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", credit, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ch *CreditHandler) HandlerPostCredit(c *fiber.Ctx) error {
	var params models.CreateCredit
	if err := readJSON(c, &params); err != nil {
		badRequestError(c)
		return err
	}

	cred, err := models.NewCreditFromParams(&params)
	if err != nil {
		badRequestError(c)
		return err
	}

	v := validator.New()
	if models.ValidateCredit(v, cred); !v.Valid() {
		err := failedValidationResponse(c, v.Errors)
		return err
	}

	credResp, err := ch.Store.InsertCredit(c.Context(), nil, cred)
	if err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", credResp, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ch *CreditHandler) HandlerUpdateCredit(c *fiber.Ctx) error {
	var params models.UpdateCredit

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

	credResp, err := ch.Store.UpdateCredit(c.Context(), nil, id, &params)
	if err != nil {
		internalServerError(c)
		return err
	}

	// TODO: Standardize messages
	err = writeJSON(c, http.StatusOK, "updated successfully", credResp, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (ch *CreditHandler) HandlerDeleteCredit(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	if err := ch.Store.DeleteCreditByID(c.Context(), nil, id); err != nil {
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
