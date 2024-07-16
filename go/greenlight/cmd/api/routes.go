package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(a.notFoundError)
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedError)

	router.HandlerFunc(http.MethodGet, "/v1/health", a.healthHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", a.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", a.getMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", a.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", a.deleteMovieHandler)

	return router
}
