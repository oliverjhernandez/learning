package api

import (
	"context"

	"finance/db"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionStore db.TransactionStore
}

func (th *TransactionHandler) HandlerTransactions(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"data": "transaction"})
}

func (th *TransactionHandler) HandlerTransaction(c *fiber.Ctx) error {
	tx, err := th.transactionStore.GetTransactionByID(context.Background(), c.Params("id"))
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
