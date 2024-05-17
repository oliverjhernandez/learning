package api

import (
	"finance/db"
	"finance/types"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionStore db.TransactionStore
}

func (th *TransactionHandler) HandlerPostTransaction(c *fiber.Ctx) error {
	var params types.TransactionParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.Validate(); err != nil {
		return err
	}

	tx, err := types.NewTransactionFromParams(params)
	if err != nil {
		return err
	}
	res, err := th.transactionStore.InsertTransaction(c.Context(), tx)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (th *TransactionHandler) HandlerGetTransactions(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"data": "many transactions"})
}

func (th *TransactionHandler) HandlerGetTransaction(c *fiber.Ctx) error {
	tx, err := th.transactionStore.GetTransactionByID(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(tx)
}

func NewTransactionHandler(transactionStore db.TransactionStore) *TransactionHandler {
	return &TransactionHandler{
		transactionStore: transactionStore,
	}
}
