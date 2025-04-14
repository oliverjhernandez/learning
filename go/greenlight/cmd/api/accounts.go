package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"greenlight/internal/data"
	"greenlight/internal/validator"
)

func (app *application) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string        `json:"title"`
		UserID   int           `json:"user_id"`
		Entity   data.Entity   `json:"entity"`
		Currency data.Currency `json:"currency"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account := &data.Account{
		Title:     input.Title,
		UserID:    input.UserID,
		Entity:    input.Entity,
		Currency:  input.Currency,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	v := validator.New()

	if data.ValidateAccount(v, account); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Accounts.Insert(account)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", account.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"account": account}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account, err := app.models.Accounts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listAccountsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string
		UserID   int
		Entity   data.Entity
		Currency data.Currency
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.UserID = app.readInt(qs, "user_id", 0, v)
	input.Entity = data.Entity(app.readInt(qs, "entity", 0, v))
	input.Currency = data.Currency(app.readInt(qs, "currency", 0, v))

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SorSafeList = []string{"id", "title", "user_id", "entity", "currency", "-id", "-title", "-user_id", "-entity", "-currency"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	accounts, metadata, err := app.models.Accounts.GetAll(
		input.Title,
		input.UserID,
		input.Entity,
		input.Currency,
		input.Filters,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"accounts": accounts, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account, err := app.models.Accounts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title    *string
		UserID   *int
		Entity   *data.Entity
		Currency *data.Currency
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		account.Title = *input.Title
	}

	if input.Entity != nil {
		account.Entity = *input.Entity
	}

	if input.Currency != nil {
		account.Currency = *input.Currency
	}

	if input.UserID != nil {
		account.UserID = *input.UserID
	}

	v := validator.New()

	if data.ValidateAccount(v, account); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Accounts.Update(account)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Accounts.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "account successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
