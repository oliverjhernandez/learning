package api

import (
	"casita/internal/db"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(stores *db.Store, app *chi.Mux) {
	var (
		txHandler      = NewTransactionHandler(stores)
		userHandler    = NewUserHandler(stores)
		creditHandler  = NewCreditHandler(stores)
		accountHandler = NewAccountHandler(stores)
		authHandler    = NewAuthHandler(stores)
	)

	app.Route("/v1", func(r chi.Router) {
		// Transaction CRUD Endpoints
		app.Route("/transactions", func(r chi.Router) {
			r.Get("/", txHandler.HandlerGetTransactions)
			r.Post("/", txHandler.HandlerPostTransaction)

			app.Route("/{txnID}", func(r chi.Router) {
				r.Get("/", txHandler.HandlerGetTransaction)
				r.Delete("/", txHandler.HandlerDeleteTransaction)
				r.Patch("/", txHandler.HandlerUpdateTransaction)
			})
		})

		// Credit CRUD Endpoints
		app.Route("/credits", func(r chi.Router) {
			r.Get("/", creditHandler.HandlerGetCredits)
			r.Post("/", creditHandler.HandlerPostCredit)

			app.Route("/{creditID}", func(r chi.Router) {
				r.Get("/:id", creditHandler.HandlerGetCredit)
				r.Delete("/:id", creditHandler.HandlerDeleteCredit)
				r.Patch("/:id", creditHandler.HandlerUpdateCredit)
			})
		})

		// User CRUD Endpoints
		app.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.HandlerGetUsers)
			r.Post("/", userHandler.HandlerPostUser)

			app.Route("/{userID}", func(r chi.Router) {
				r.Get("/:id", userHandler.HandlerGetUser)
				r.Delete("/:id", userHandler.HandlerDeleteUser)
				r.Patch("/:id", userHandler.HandlerUpdateUser)
			})
		})

		// Account CRUD Endpoints
		app.Route("/accounts", func(r chi.Router) {
			r.Get("/", accountHandler.HandlerGetAccounts)
			r.Post("/", accountHandler.HandlerPostAccount)

			app.Route("/{accID}", func(r chi.Router) {
				r.Get("/", accountHandler.HandlerGetAccount)
				r.Delete("/", accountHandler.HandlerDeleteAccount)
				r.Patch("/", accountHandler.HandlerUpdateAccount)
			})
		})
	})

	app.Route("/api", func(r chi.Router) {
		// Auth
		r.Post("/auth", authHandler.HandleAuthenticate)
	})
}
