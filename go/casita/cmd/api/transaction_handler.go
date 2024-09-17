package api

import (
	"errors"
	"net/http"
	"strconv"

	"casita/internal/db"
	"casita/internal/models"
	"casita/internal/validator"

	"github.com/go-chi/chi/v5"
)

type TransactionHandler struct {
	Store db.TransactionStore
}

func NewTransactionHandler(s *db.Store) *TransactionHandler {
	return &TransactionHandler{
		Store: s.TxnStore,
	}
}

func (th *TransactionHandler) HandlerPostTransaction(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params models.CreateTransaction
	if err := readJSON(c, &params); err != nil {
		badRequestError(err)
		return
	}

	txn, err := models.NewTransactionFromParams(&params)
	if err != nil {
		badRequestError(err)
		return
	}

	v := validator.New()
	if models.ValidateTransaction(v, txn); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	tran, err := th.Store.InsertTransaction(c, nil, txn)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(c, http.StatusOK, "resource created successfully", tran, nil, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (th *TransactionHandler) HandlerGetTransactions(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	input := models.ListTransactions{}
	v := validator.New()

	input.Concept = readString(c, "concept", "")
	input.Description = readString(c, "description", "")
	input.Relevance = models.Relevance(readInt(c, "relevance", 0, v))
	input.Value = int32(readInt(c, "value", -1, v))

	input.Page = readInt(c, "page", 1, v)
	input.PageSize = readInt(c, "page_size", 10, v)
	input.Sort = readString(c, "sort", "value")
	input.SortSafeList = []string{"value", "-value", "concept", "-concept", "relevance", "-relevance", "day", "-day", "month", "-month"}

	if models.ValidateFilters(v, input.Filters); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	if !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	txns, metadata, err := th.Store.GetAllTransactions(c, nil, input.Concept, input.Value, input.Description, input.Filters)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(c, http.StatusOK, "got you", &txns, &metadata, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (th *TransactionHandler) HandlerGetTransaction(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	txn, err := th.Store.GetTransactionByID(c, nil, id)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(c, http.StatusOK, "got you", &txn, nil, "")
	if err != nil {
		internalServerError(err)
	}

	return
}

func (th *TransactionHandler) HandlerUpdateTransaction(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	var params models.UpdateTransaction

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	if err := readJSON(c, &params); err != nil {
		badRequestError(err)
		return
	}

	tran, err := th.Store.UpdateTransaction(c, nil, id, &params)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(c, http.StatusOK, "resource updated successfully", &tran, nil, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (th *TransactionHandler) HandlerDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	if err := th.Store.DeleteTransactionByID(c, nil, id); err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(c, http.StatusOK, "resorce deleted successfully", nil, nil, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}
