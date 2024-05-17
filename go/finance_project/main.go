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
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")

	app := fiber.New(config)
	appv1 := app.Group("/v1")

	appv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Stores
	txStore := db.NewMongoTransactionStore(client)
	// Handlers
	txHandler := api.NewTransactionHandler(txStore)

	appv1.Get("/transaction", txHandler.HandlerGetTransactions)
	appv1.Get("/transaction/:id", txHandler.HandlerGetTransaction)
	appv1.Post("/transaction", txHandler.HandlerPostTransaction)

	app.Listen(*listenAddr)
}
