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

type UserHandler struct {
	Store db.UserStore
}

func NewUserHandler(s *db.Store) *UserHandler {
	return &UserHandler{
		Store: s.UserStore,
	}
}

func (uh *UserHandler) HandlerPostUser(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params models.CreateUser
	if err := readJSON(r, &params); err != nil {
		badRequestError(err)
		return
	}

	user, err := models.NewUserFromParams(&params)
	if err != nil {
		badRequestError(err)
		return
	}

	v := validator.New()
	if models.ValidateUser(v, user); !v.Valid() {
		unprocessableEntityError(errors.New("unprocessableEntityError"))
		return
	}

	userResp, err := uh.Store.InsertUser(c, nil, user)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(w, userResp)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (uh *UserHandler) HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	users, err := uh.Store.GetAllUsers(c, nil)
	if err != nil {
		notFoundError(err)
		return
	}

	err = writeJSON(w, users)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (uh *UserHandler) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	user, err := uh.Store.GetUserByID(c, nil, id)
	if err != nil {
		badRequestError(err)
		return
	}

	err = writeJSON(w, user)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (uh *UserHandler) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params models.UpdateUser

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

	userResp, err := uh.Store.UpdateUser(c, nil, id, &params)
	if err != nil {
		internalServerError(err)
		return
	}

	err = writeJSON(w, userResp)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func (uh *UserHandler) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		badRequestError(err)
		return
	}

	if err := uh.Store.DeleteUserByID(c, nil, id); err != nil {
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
