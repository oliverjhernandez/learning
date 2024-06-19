package api

import (
	"strconv"

	"finance/db"
	"finance/models"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Store *db.PGTransactionStore // TODO: This should be a more general type of store
}

func (th *TransactionHandler) HandlerGetTransactions(c *fiber.Ctx) error {
	txs, err := th.Store.GetAllTransactions()
	if err != nil {
		return ErrNotFound()
	}
	return c.JSON(&txs)
}

func (th *TransactionHandler) HandlerGetTransaction(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	tx, err := th.Store.GetTransactionByID(id)
	if err != nil {
		return ErrNotFound()
	}

	return c.JSON(tx)
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

	tx := models.NewTransactionFromParams(params)
	res, err := th.Store.InsertTransaction(tx)
	if err != nil {
		return err
	}
	return c.JSON(res)
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

	if err = th.Store.UpdateTransaction(id, &params); err != nil {
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

	if err := th.Store.DeleteTransactionByID(id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "deleted"})
}

func NewTransactionHandler(s *db.PGTransactionStore) *TransactionHandler {
	return &TransactionHandler{
		Store: s,
	}
}
