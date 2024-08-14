package api

import (
	"fmt"
	"net/http"
	"strconv"

	"casita/internal/db"
	"casita/internal/models"
	"casita/internal/validator"

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
	if err := readJSON(c, &params); err != nil {
		badRequestError(c)
		return err
	}

	txn, err := models.NewTransactionFromParams(&params)
	if err != nil {
		badRequestError(c)
		return err
	}

	v := validator.New()
	if models.ValidateTransaction(v, txn); !v.Valid() {
		err := failedValidationResponse(c, v.Errors)
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
	input := models.ListTransactions{}
	v := validator.New()

	input.Concept = readString(c, "concept", "")
	input.Description = readString(c, "description", "")
	input.Relevance = models.Relevance(readInt(c, "relevance", 0, v))
	input.Value = int32(readInt(c, "value", -1, v))

	input.Page = readInt(c, "page", 1, v)
	input.PageSize = readInt(c, "page_size", 10, v)
	input.Sort = readString(c, "sort", "value")
	input.SortSafeList = []string{"value", "-value", "relevance", "-relevance", "day", "-day", "month", "-month"}

	if models.ValidateFilters(v, input.Filters); !v.Valid() {
		err := failedValidationResponse(c, v.Errors)
		return err
	}

	fmt.Printf("Input: %+v\n", input)

	if !v.Valid() {
		failedValidationResponse(c, v.Errors)
		err := failedValidationResponse(c, v.Errors)
		return err
	}

	txns, err := th.Store.GetAllTransactions(c.Context(), nil, input.Concept, input.Value, input.Description, input.Filters)
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

	if err := readJSON(c, &params); err != nil {
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
