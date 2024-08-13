package api

import (
	"casita/internal/db"

	"github.com/gofiber/fiber/v2"
)

func Routes(stores *db.Store, app *fiber.App) {
	var (
		txHandler      = NewTransactionHandler(stores)
		userHandler    = NewUserHandler(stores)
		creditHandler  = NewCreditHandler(stores)
		accountHandler = NewAccountHandler(stores)
		authHandler    = NewAuthHandler(stores)

		appv1 = app.Group("/v1")
		// appv1 = fiberApp.Group("/v1", api.JWTAuthentication(stores.UserStore))
		api = app.Group("/api")
		// admin = appv1.Group("/admin", api.AdminAuth)
	)

	// Auth
	api.Post("/auth", authHandler.HandleAuthenticate)

	// Transaction CRUD Endpoints
	appv1.Get("/transaction", txHandler.HandlerGetTransactions)
	appv1.Get("/transaction/:id", txHandler.HandlerGetTransaction)
	appv1.Post("/transaction", txHandler.HandlerPostTransaction)
	appv1.Delete("/transaction/:id", txHandler.HandlerDeleteTransaction)
	appv1.Patch("/transaction/:id", txHandler.HandlerUpdateTransaction)

	// Credit CRUD Endpoints
	appv1.Get("/credit", creditHandler.HandlerGetCredits)
	appv1.Get("/credit/:id", creditHandler.HandlerGetCredit)
	appv1.Post("/credit", creditHandler.HandlerPostCredit)
	appv1.Delete("/credit/:id", creditHandler.HandlerDeleteCredit)
	appv1.Patch("/credit/:id", creditHandler.HandlerUpdateCredit)

	// User CRUD Endpoints
	appv1.Get("/user", userHandler.HandlerGetUsers)
	appv1.Get("/user/:id", userHandler.HandlerGetUser)
	appv1.Post("/user", userHandler.HandlerPostUser)
	appv1.Delete("/user/:id", userHandler.HandlerDeleteUser)
	appv1.Patch("/user/:id", userHandler.HandlerUpdateUser)

	// Account CRUD Endpoints
	appv1.Get("/account", accountHandler.HandlerGetAccounts)
	appv1.Get("/account/:id", accountHandler.HandlerGetAccount)
	appv1.Post("/account", accountHandler.HandlerPostAccount)
	appv1.Delete("/account/:id", accountHandler.HandlerDeleteAccount)
	appv1.Patch("/account/:id", accountHandler.HandlerUpdateAccount)
}
