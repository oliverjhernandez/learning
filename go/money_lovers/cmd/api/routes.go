package api

import (
	"context"
	"net/http"

	"money_lovers/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func InitializeRoutes(app *chi.Mux, dbCfg db.DBCfg, logger *httplog.Logger) {
	client, err := db.ConnectSQL(dbCfg.DB)
	if err != nil {
		logger.Log(context.TODO(), 8, err.Error())
	}

	logger.Log(context.TODO(), 0, "database connection pool stablished")
	// defer func() {
	// if err := client.Close(); err != nil {
	//   logger.Log(context.TODO(), 8, "error closing connection to db: "+err.Error())
	// }
	// }()

	stores := &db.Store{
		DB:           client,
		UserStore:    db.NewPGUserStore(client),
		TxnStore:     db.NewPGTransactionStore(client),
		CreditStore:  db.NewPGCreditStore(client),
		AccountStore: db.NewPGAccountStore(client),
	}

	var (
		txHandler      = NewTransactionHandler(stores)
		userHandler    = NewUserHandler(stores, logger)
		creditHandler  = NewCreditHandler(stores)
		accountHandler = NewAccountHandler(stores)
		authHandler    = NewAuthHandler(stores)
	)

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	app.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// INFO: Panic endpoint for testing
	app.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	app.Route("/v1", func(r chi.Router) {
		// Transaction CRUD Endpoints
		r.Route("/transactions", func(r chi.Router) {
			r.Get("/", txHandler.HandlerGetTransactions)
			r.Post("/", txHandler.HandlerPostTransaction)

			r.Route("/{txnID}", func(r chi.Router) {
				r.Get("/", txHandler.HandlerGetTransaction)
				r.Delete("/", txHandler.HandlerDeleteTransaction)
				r.Patch("/", txHandler.HandlerUpdateTransaction)
			})
		})

		// Credit CRUD Endpoints
		r.Route("/credits", func(r chi.Router) {
			r.Get("/", creditHandler.HandlerGetCredits)
			r.Post("/", creditHandler.HandlerPostCredit)

			r.Route("/{credID}", func(r chi.Router) {
				r.Get("/", creditHandler.HandlerGetCredit)
				r.Delete("/", creditHandler.HandlerDeleteCredit)
				r.Patch("/", creditHandler.HandlerUpdateCredit)
			})
		})

		// User CRUD Endpoints
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.HandlerGetUsers)
			r.Post("/", userHandler.HandlerPostUser)

			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", userHandler.HandlerGetUser)
				r.Delete("/", userHandler.HandlerDeleteUser)
				r.Patch("/", userHandler.HandlerUpdateUser)
			})
		})

		// Account CRUD Endpoints
		r.Route("/accounts", func(r chi.Router) {
			r.Get("/", accountHandler.HandlerGetAccounts)
			r.Post("/", accountHandler.HandlerPostAccount)

			r.Route("/{accID}", func(r chi.Router) {
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
