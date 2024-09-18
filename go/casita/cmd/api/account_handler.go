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

type AccountHandler struct {
	Store db.AccountStore
}

func NewAccountHandler(s *db.Store) *AccountHandler {
	return &AccountHandler{
		Store: s.AccountStore,
	}
}

var params *models.CreateAccount

func (ah *AccountHandler) HandlerPostAccount(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params models.CreateAccount
	if err := readJSON(r, &params); err != nil {
		badRequestError(err)
		return
	}

	acc, err := models.NewAccountFromParams(&params)
	if err != nil {
		badRequestError(err)
		return
	}

	v := validator.New()
	if models.ValidateAccount(v, acc); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	acc, err = ah.Store.InsertAccount(c, nil, acc)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(w, acc)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ah *AccountHandler) HandlerGetAccounts(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	accs, err := ah.Store.GetAllAccounts(c, nil)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(w, accs)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ah *AccountHandler) HandlerGetAccount(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	acc, err := ah.Store.GetAccountByID(c, nil, id)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(w, acc)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ah *AccountHandler) HandlerUpdateAccount(w http.ResponseWriter, r *http.Request) {
	var params models.UpdateAccount

	c := r.Context()
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

	acc, err := ah.Store.UpdateAccount(c, nil, id, &params)
	if err != nil {
		editConflictError(err)
		return
	}

	err = writeJSON(w, acc)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ah *AccountHandler) HandlerDeleteAccount(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	if err := ah.Store.DeleteAccountByID(c, nil, id); err != nil {
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
