package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(a.notFoundError)
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedError)

	router.HandlerFunc("GET", "/v1/health", a.healthHandler)
	router.HandlerFunc("POST", "/v1/movie", a.createMovieHandler)
	router.HandlerFunc("GET", "/v1/movie/:id", a.getMovieHandler)

	return router
}
