package api

import (
	"strconv"

	"finance/db"
	"finance/models"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Store db.TransactionStore
}

func NewTransactionHandler(s *db.Store) *TransactionHandler {
	return &TransactionHandler{
		Store: s.TxnStore,
	}
}

func (th *TransactionHandler) HandlerPostTransaction(c *fiber.Ctx) error {
	var params models.CreateTransaction
	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	// TODO: There should be some validation of the data coming in
	//
	// if err := params.Validate(); err != nil {
	// 	return ErrInvalidParams()
	// }

	txn := models.NewTransactionFromParams(params)
	res, err := th.Store.InsertTransaction(c.Context(), nil, txn)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (th *TransactionHandler) HandlerGetTransactions(c *fiber.Ctx) error {
	txns, err := th.Store.GetAllTransactions(c.Context(), nil)
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(&txns)
}

func (th *TransactionHandler) HandlerGetTransaction(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	txn, err := th.Store.GetTransactionByID(c.Context(), nil, id)
	if err != nil {
		return ErrNotFound()
	}

	return c.JSON(txn)
}

func (th *TransactionHandler) HandlerUpdateTransaction(c *fiber.Ctx) error {
	var params models.UpdateTransaction

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&params); err != nil {
		return ErrInvalidReqBody()
	}

	if err = th.Store.UpdateTransaction(c.Context(), nil, id, &params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "updated"})
}

func (th *TransactionHandler) HandlerDeleteTransaction(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	if err := th.Store.DeleteTransactionByID(c.Context(), nil, id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "deleted"})
}
