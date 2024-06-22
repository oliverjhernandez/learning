package main

import (
	"flag"
	"fmt"

	"finance/api"
	"finance/db"

	"github.com/gofiber/fiber/v2"
)

const (
	dburi = "mongodb://localhost:27017"
	pg_db = ""
)

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	store, client, err := db.NewStore()
	if err != nil {
		fmt.Printf("error loading db store")
	}
	defer client.Close()

	app := fiber.New(config)
	appv1 := app.Group("/v1")

	appv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	var (
		listenAddr = flag.String("listenAddr", ":3000", "The listen address of the API server")
		// store, _   = db.NewStore(client)

		// Handlers
		txHandler     = api.NewTransactionHandler(store)
		userHandler   = api.NewUserHandler(store)
		creditHandler = api.NewCreditHandler(store)
	)

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

	app.Listen(*listenAddr)
}
