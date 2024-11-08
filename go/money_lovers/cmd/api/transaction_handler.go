package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"money_lovers/internal/db"
	"money_lovers/internal/validator"

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
	var params db.CreateTransaction
	if err := readJSON(r, &params); err != nil {
		badRequestError(err)
		return
	}

	txn, err := db.NewTransactionFromParams(&params)
	if err != nil {
		badRequestError(err)
		return
	}

	v := validator.New()
	if db.ValidateTransaction(v, txn); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	tran, err := th.Store.InsertTransaction(c, nil, txn)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(w, tran)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (th *TransactionHandler) HandlerGetTransactions(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	i := db.GetTransactions{}
	v := validator.New()

	i.Concept = readString(r, "concept", "")
	i.Description = readString(r, "description", "")
	i.Relevance = db.Relevance(readInt(r, "relevance", 0, v))
	i.Value = int32(readInt(r, "value", -1, v))

	i.Page = readInt(r, "page", 1, v)
	i.PageSize = readInt(r, "page_size", 10, v)
	i.Sort = readString(r, "sort", "value")
	i.SortSafeList = []string{"value", "-value", "concept", "-concept", "relevance", "-relevance", "day", "-day", "month", "-month"}

	if db.ValidateFilters(v, i.Filters); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJSON(w, map[string]string{"error": "unprocessableEntityError"})
		return
	}

	if !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJSON(w, map[string]string{"error": "unprocessableEntityError"})
		return
	}

	txns, _, err := th.Store.GetAllTransactions(c, nil, i.Concept, i.Value, i.Description, i.Filters)
	if err != nil {
		notFoundError(err)
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	fmt.Println(txns)

	w.WriteHeader(http.StatusOK)
	err = writeJSON(w, &txns)
	if err != nil {
		internalServerError(err)
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
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

	err = writeJSON(w, &txn)
	if err != nil {
		internalServerError(err)
	}

	return
}

func (th *TransactionHandler) HandlerUpdateTransaction(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	var params db.UpdateTransaction

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	if err := readJSON(r, &params); err != nil {
		badRequestError(err)
		return
	}

	tran, err := th.Store.UpdateTransaction(c, nil, id, &params)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(w, &tran)
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

	err = writeJSON(w, nil)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}
