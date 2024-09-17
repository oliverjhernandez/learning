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

	err = writeJSON(c, http.StatusOK, "got you", credits, nil, "")
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

	err = writeJSON(c, http.StatusOK, "got you", credit, nil, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ch *CreditHandler) HandlerPostCredit(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params models.CreateCredit
	if err := readJSON(c, &params); err != nil {
		badRequestError(err)
		return
	}

	cred, err := models.NewCreditFromParams(&params)
	if err != nil {
		badRequestError((err))
		return
	}

	v := validator.New()
	if models.ValidateCredit(v, cred); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	credResp, err := ch.Store.InsertCredit(c, nil, cred)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(c, http.StatusOK, "got you", credResp, nil, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (ch *CreditHandler) HandlerUpdateCredit(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params models.UpdateCredit

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

	credResp, err := ch.Store.UpdateCredit(c, nil, id, &params)
	if err != nil {
		internalServerError(err)
		return
	}

	// TODO: Standardize messages
	err = writeJSON(c, http.StatusOK, "updated successfully", credResp, nil, "")
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

	err = writeJSON(c, http.StatusOK, "resource deleted", nil, nil, "")
	if err != nil {
		internalServerError(err)
		return
	}

	return
}
