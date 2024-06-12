package main

import (
	"context"
	"flag"

	"finance/api"
	"finance/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi = "mongodb://localhost:27017"
)

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}

	app := fiber.New(config)
	appv1 := app.Group("/v1")

	appv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	var (
		listenAddr = flag.String("listenAddr", ":3000", "The listen address of the API server")
		// Stores
		txStore   = db.NewMongoTransactionStore(client, db.DBNAME)
		userStore = db.NewMongoUserStore(client, db.DBNAME)
		store     = &db.Store{
			User: userStore,
			Tx:   txStore,
		}
		// Handlers
		txHandler   = api.NewTransactionHandler(store)
		userHandler = api.NewUserHandler(store)
	)

	// Transaction CRUD Endpoints
	appv1.Get("/transaction", txHandler.HandlerGetTransactions)
	appv1.Get("/transaction/:id", txHandler.HandlerGetTransaction)
	appv1.Post("/transaction", txHandler.HandlerPostTransaction)
	appv1.Delete("/transaction/:id", txHandler.HandlerDeleteTransaction)
	appv1.Patch("/transaction/:id", txHandler.HandlerUpdateTransaction)

	// User CRUD Endpoints
	appv1.Get("/user", userHandler.HandlerGetUsers)
	appv1.Get("/user/:id", userHandler.HandlerGetUser)
	appv1.Post("/user", userHandler.HandlerPostUser)
	appv1.Delete("/user/:id", userHandler.HandlerDeleteUser)
	appv1.Patch("/user/:id", userHandler.HandlerUpdateUser)

	app.Listen(*listenAddr)
}
