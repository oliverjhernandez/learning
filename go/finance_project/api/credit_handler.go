package api

import (
	"finance/db"
	"finance/types"

	"github.com/gofiber/fiber/v2"
)

type CreditHandler struct {
	Store *db.Store
}

func NewCreditHandler(store *db.Store) *CreditHandler {
	return &CreditHandler{
		Store: store,
	}
}

func (ch *CreditHandler) HandlerGetCredits(c *fiber.Ctx) error {
	credits, err := ch.Store.Credit.GetCredits(c.Context())
	if err != nil {
		return ErrNotFound()
	}

	return c.JSON(&credits)
}

func (ch *CreditHandler) HandlerGetCredit(c *fiber.Ctx, id string) error {
	credit, err := ch.Store.Credit.GetCreditByID(c.Context(), id)
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(&credit)
}

func (ch *CreditHandler) HandlerPostCredit(c *fiber.Ctx, credit *types.Credit) error {
	var params *types.CreateCreditParams
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidParams()
	}

	cred, err := types.NewCreditFromParams(params)
	if err != nil {
		return err
	}

	_, err = ch.Store.Credit.InsertCredit(c.Context(), cred)
	if err != nil {
		return err
	}

	return c.JSON(cred)
}
