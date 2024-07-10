package main

import (
	"fmt"
	"net/http"
	"time"

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

	fmt.Fprintf(w, "%+v\n", input)
}

func (a *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDFromParams(r)
	if err != nil {
		a.notFoundError(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "The Silence of the Lambs",
		Year:      1996,
		Runtime:   102,
		Genres:    []string{"suspense", "thriller"},
		Version:   1,
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		a.internalServerError(w, r)
	}
}
