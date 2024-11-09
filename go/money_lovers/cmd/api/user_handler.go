package api

import (
	"errors"
	"net/http"
	"strconv"

	"money_lovers/internal/db"
	"money_lovers/internal/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

type UserHandler struct {
	Store  db.UserStore
	Logger *httplog.Logger
}

func NewUserHandler(s *db.Store, logger *httplog.Logger) *UserHandler {
	return &UserHandler{
		Store:  s.UserStore,
		Logger: logger,
	}
}

func (uh *UserHandler) HandlerPostUser(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params db.CreateUser
	if err := readJSON(r, &params); err != nil {
		se := badRequestError(err)
		writeErrorResponse(w, se)
		return
	}

	user, err := db.NewUserFromParams(&params)
	if err != nil {
		se := badRequestError(err)
		writeErrorResponse(w, se)
		return
	}

	v := validator.New()
	if db.ValidateUser(v, user); !v.Valid() {
		se := unprocessableEntityError(errors.New("unprocessableEntityError"))
		writeErrorResponse(w, se)
		return
	}

	userResp, err := uh.Store.InsertUser(c, nil, user)
	if err != nil {
		se := conflictError(err)
		writeErrorResponse(w, se)
		return
	}

	err = writeJSON(w, userResp)
	if err != nil {
		se := internalServerError(err)
		writeErrorResponse(w, se)
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
	var params db.UpdateUser

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
