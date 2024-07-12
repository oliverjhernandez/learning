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
	router.HandlerFunc(http.MethodPost, "/v1/movie", a.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", a.getMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movie/:id", a.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movie/:id", a.deleteMovieHandler)

	return router
}
