package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight/internal/data"
)

func (a *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
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
