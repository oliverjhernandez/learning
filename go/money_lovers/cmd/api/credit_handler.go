package api

import (
	"errors"
	"net/http"
	"strconv"

	"money_lovers/internal/db"
	"money_lovers/internal/validator"

	"github.com/go-chi/chi/v5"
)

type CreditHandler struct {
	Store db.CreditStore
}

func NewCreditHandler(s *db.Store) *CreditHandler {
	return &CreditHandler{
		Store: s.CreditStore,
	}
}

func (ch *CreditHandler) HandlerGetCredits(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	credits, err := ch.Store.GetAllCredits(c, nil)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(w, credits)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ch *CreditHandler) HandlerGetCredit(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	credit, err := ch.Store.GetCreditByID(c, nil, id)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(w, credit)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ch *CreditHandler) HandlerPostCredit(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params db.CreateCredit
	if err := readJSON(r, &params); err != nil {
		badRequestError(err)
		return
	}

	cred, err := db.NewCreditFromParams(&params)
	if err != nil {
		badRequestError((err))
		return
	}

	v := validator.New()
	if db.ValidateCredit(v, cred); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	credResp, err := ch.Store.InsertCredit(c, nil, cred)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(w, credResp)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ch *CreditHandler) HandlerUpdateCredit(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params db.UpdateCredit

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

	credResp, err := ch.Store.UpdateCredit(c, nil, id, &params)
	if err != nil {
		internalServerError(err)
		return
	}

	// TODO: Standardize messages
	err = writeJSON(w, credResp)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ch *CreditHandler) HandlerDeleteCredit(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	if err := ch.Store.DeleteCreditByID(c, nil, id); err != nil {
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
