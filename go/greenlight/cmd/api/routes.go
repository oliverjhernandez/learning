package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// TODO: Move this to standard library
	// router := httprouter.New()
	// router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// healthcheck
	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	// movies
	mux.HandleFunc("POST /v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
	mux.HandleFunc("GET /v1/movies/{id}", app.requirePermission("movies:read", app.showMovieHandler))
	mux.HandleFunc("GET /v1/movies", app.requirePermission("movies:read", app.listMoviesHandler))
	mux.HandleFunc("PATCH /v1/movies/{id}", app.requirePermission("movies:write", app.updateMovieHandler))
	mux.HandleFunc("DELETE /v1/movies/{id}", app.requirePermission("movies:write", app.deleteMovieHandler))

	// accounts
	mux.HandleFunc("POST /v1/accounts", app.requirePermission("accounts:write", app.createAccountHandler))
	mux.HandleFunc("GET /v1/accounts/{id}", app.requirePermission("accounts:read", app.showAccountHandler))
	mux.HandleFunc("GET /v1/accounts", app.requirePermission("accounts:read", app.listAccountsHandler))
	mux.HandleFunc("PATCH /v1/accounts/{id}", app.requirePermission("accounts:write", app.updateAccountHandler))
	mux.HandleFunc("DELETE /v1/accounts/{id}", app.requirePermission("accounts:write", app.deleteAccountHandler))

	// users
	mux.HandleFunc("POST /v1/users", app.registerUserHandler)
	mux.HandleFunc("PUT /v1/users/activated", app.activateUserHandler)

	mux.HandleFunc("POST /v1/tokens/activation", app.createActivationTokenHandler)
	mux.HandleFunc("POST /v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(
		app.recoverPanic(
			app.enableCORS(
				app.rateLimit(
					app.authenticate(mux),
				),
			),
		),
	)
}
