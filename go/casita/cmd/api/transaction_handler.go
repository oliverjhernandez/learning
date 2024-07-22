package api

import (
	"net/http"
	"strconv"

	"casita/internal/db"
	"casita/internal/models"

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
		badRequestError(c)
		return err
	}

	// TODO: There should be some validation of the data coming in
	//
	// if err := params.Validate(); err != nil {
	// 	return ErrInvalidParams()
	// }

	txn, err := models.NewTransactionFromParams(&params)
	if err != nil {
		badRequestError(c)
		return err
	}
	tran, err := th.Store.InsertTransaction(c.Context(), nil, txn)
	if err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", tran, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (th *TransactionHandler) HandlerGetTransactions(c *fiber.Ctx) error {
	txns, err := th.Store.GetAllTransactions(c.Context(), nil)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", &txns, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (th *TransactionHandler) HandlerGetTransaction(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	txn, err := th.Store.GetTransactionByID(c.Context(), nil, id)
	if err != nil {
		notFoundError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "got you", &txn, "")
	if err != nil {
		internalServerError(c)
	}

	return nil
}

func (th *TransactionHandler) HandlerUpdateTransaction(c *fiber.Ctx) error {
	var params models.UpdateTransaction

	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	if err := c.BodyParser(&params); err != nil {
		badRequestError(c)
		return err
	}

	tran, err := th.Store.UpdateTransaction(c.Context(), nil, id, &params)
	if err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resource updated successfully", &tran, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}

func (th *TransactionHandler) HandlerDeleteTransaction(c *fiber.Ctx) error {
	strID := c.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(c)
		return err
	}

	if err := th.Store.DeleteTransactionByID(c.Context(), nil, id); err != nil {
		internalServerError(c)
		return err
	}

	err = writeJSON(c, http.StatusOK, "resorce deleted successfully", nil, "")
	if err != nil {
		internalServerError(c)
		return err
	}

	return nil
}
