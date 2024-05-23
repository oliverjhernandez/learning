package api

import (
	"finance/db"
	"finance/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionHandler struct {
	transactionStore db.TransactionStore
}

func (th *TransactionHandler) HandlerGetTransactions(c *fiber.Ctx) error {
	txs, err := th.transactionStore.GetTransactions(c.Context())
	if err != nil {
		return nil
	}
	return c.JSON(&txs)
}

func (th *TransactionHandler) HandlerGetTransaction(c *fiber.Ctx) error {
	tx, err := th.transactionStore.GetTransactionByID(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(tx)
}

func (th *TransactionHandler) HandlerPostTransaction(c *fiber.Ctx) error {
	var params types.CreateTransactionParams
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

func (th *TransactionHandler) HandlerUpdateTransaction(c *fiber.Ctx) error {
	var (
		params types.UpdateTransactionParams
		userID = c.Params("id")
	)

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	if err = th.transactionStore.UpdateTransaction(c.Context(), filter, &params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "updated"})
}

func (th *TransactionHandler) HandlerDeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := th.transactionStore.DeleteTransaction(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "deleted"})
}

func NewTransactionHandler(transactionStore db.TransactionStore) *TransactionHandler {
	return &TransactionHandler{
		transactionStore: transactionStore,
	}
}
