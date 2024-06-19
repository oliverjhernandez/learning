package api

import (
	"strconv"

	"finance/db"
	"finance/models"

	"github.com/gofiber/fiber/v2"
)

type CreditHandler struct {
	Store *db.PGCreditStore
}

func NewCreditHandler(store *db.PGCreditStore) *CreditHandler {
	return &CreditHandler{
		Store: store,
	}
}

func (ch *CreditHandler) HandlerGetCredits(c *fiber.Ctx) error {
	credits, err := ch.Store.GetAllCredits()
	if err != nil {
		return ErrNotFound()
	}

	return c.JSON(&credits)
}

func (ch *CreditHandler) HandlerGetCredit(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	credit, err := ch.Store.GetCreditByID(id)
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(&credit)
}

func (ch *CreditHandler) HandlerPostCredit(c *fiber.Ctx) error {
	var params *models.CreateCreditParams
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidParams()
	}

	cred := models.NewCreditFromParams(params)

	_, err := ch.Store.InsertCredit(cred)
	if err != nil {
		return err
	}

	return c.JSON(cred)
}

func (ch *CreditHandler) HandlerUpdateCredit(c *fiber.Ctx) error {
	var params models.UpdateCreditParams

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	ch.Store.UpdateCredit(id, &params)

	return c.JSON(map[string]string{"msg": "updated"})
}

func (ch *CreditHandler) HandlerDeleteCredit(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	ch.Store.DeleteCreditByID(id)

	return c.JSON(map[string]string{"msg": "deleted"})
}
