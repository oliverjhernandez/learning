package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"greenlight/internal/data"
	"greenlight/internal/validator"
)

func (a *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestError(w, r)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := a.models.Movies.Insert(movie); err != nil {
		// TODO: improve error here
		a.internalServerError(w, r)
		panic(err)
	}

	headers := make(http.Header)
	headers.Set("Locations", fmt.Sprintf("/v1/movies/%d", movie.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		a.internalServerError(w, r)
	}
}

func (a *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDFromParams(r)
	if err != nil {
		a.notFoundError(w, r)
		return
	}

	movie, err := a.models.Movies.Get(int(id))
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundError(w, r)
		default:
			a.internalServerError(w, r)
		}
		return

	}

	err = a.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		a.internalServerError(w, r)
	}
}

func (a *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDFromParams(r)
	if err != nil {
		a.notFoundError(w, r)
		return
	}

	movie, err := a.models.Movies.Get(int(id))
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundError(w, r)
		default:
			a.errorResponse(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(movie.Version), 32) != r.Header.Get("X-Expected-Version") {
			a.editConflictResponse(w, r)
			return
		}
	}

	var input struct {
		Title   *string       `json:"title"`
		Year    *int32        `json:"year"`
		Runtime *data.Runtime `json:"runtime"`
		Genres  []string      `json:"genres"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestError(w, r)
		return
	}

	if input.Title != nil {
		movie.Title = *input.Title
	}

	if input.Year != nil {
		movie.Year = *input.Year
	}

	if input.Runtime != nil {
		movie.Runtime = *input.Runtime
	}

	if input.Genres != nil {
		movie.Genres = input.Genres
	}

	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.models.Movies.Update(movie)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConfilct):
			a.editConflictResponse(w, r)
		default:
			a.errorResponse(w, r, http.StatusInternalServerError, err)
		}
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		a.errorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}

func (a *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDFromParams(r)
	if err != nil {
		a.badRequestError(w, r)
		return
	}

	if err = a.models.Movies.Delete(int(id)); err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundError(w, r)
		default:
			a.internalServerError(w, r)
		}
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		a.errorResponse(w, r, http.StatusInternalServerError, err)
	}
}
