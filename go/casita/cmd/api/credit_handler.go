package api

import (
	"net/http"
	"strconv"

	"casita/db"
	"casita/models"

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
		return NewError(http.StatusNotFound, NOT_FOUND)
	}

	return c.JSON(&credits)
}

func (ch *CreditHandler) HandlerGetCredit(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	credit, err := ch.Store.GetCreditByID(c.Context(), nil, id)
	if err != nil {
		return NewError(http.StatusNotFound, NOT_FOUND)
	}
	return c.JSON(&credit)
}

func (ch *CreditHandler) HandlerPostCredit(c *fiber.Ctx) error {
	var params *models.CreateCredit
	if err := c.BodyParser(&params); err != nil {
		return NewError(http.StatusBadRequest, INVALID_PARAMETERS)
	}

	cred, err := models.NewCreditFromParams(params)
	if err != nil {
		return err
	}

	_, err = ch.Store.InsertCredit(c.Context(), nil, cred)
	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"msg": "created"})
}

func (ch *CreditHandler) HandlerUpdateCredit(c *fiber.Ctx) error {
	var params models.UpdateCredit

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&params); err != nil {
		return NewError(http.StatusBadRequest, INVALID_PARAMETERS)
	}

	if err := ch.Store.UpdateCredit(c.Context(), nil, id, &params); err != nil {
		return err
	}

	// TODO: Serialize good path responses
	return c.JSON(map[string]string{"msg": "updated"})
}

func (ch *CreditHandler) HandlerDeleteCredit(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := ch.Store.DeleteCreditByID(c.Context(), nil, id); err != nil {
		return err
	}

	return c.JSON(map[string]string{"msg": "deleted"})
}
