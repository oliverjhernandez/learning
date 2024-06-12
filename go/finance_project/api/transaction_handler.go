package api

import (
	"finance/db"
	"finance/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionHandler struct {
	Store *db.Store
}

func (th *TransactionHandler) HandlerGetTransactions(c *fiber.Ctx) error {
	txs, err := th.Store.Tx.GetTransactions(c.Context())
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(&txs)
}

func (th *TransactionHandler) HandlerGetTransaction(c *fiber.Ctx) error {
	tx, err := th.Store.Tx.GetTransactionByID(c.Context(), c.Params("id"))
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(tx)
}

func (th *TransactionHandler) HandlerPostTransaction(c *fiber.Ctx) error {
	var params types.CreateTransactionParams
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}
	if err := params.Validate(); err != nil {
		return ErrInvalidParams()
	}

	tx, err := types.NewTransactionFromParams(params)
	if err != nil {
		return err
	}
	res, err := th.Store.Tx.InsertTransaction(c.Context(), tx)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (th *TransactionHandler) HandlerUpdateTransaction(c *fiber.Ctx) error {
	var (
		params types.UpdateTransactionParams
		userID = c.Params("id")
	)

	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrInvalidID()
	}

	filter := map[string]any{"_id": oid}

	if err = th.Store.Tx.UpdateTransaction(c.Context(), filter, &params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "updated"})
}

func (th *TransactionHandler) HandlerDeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := th.Store.Tx.DeleteTransaction(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "deleted"})
}

func NewTransactionHandler(s *db.Store) *TransactionHandler {
	return &TransactionHandler{
		Store: s,
	}
}
