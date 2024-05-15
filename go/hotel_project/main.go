package main

import (
	"context"
	"flag"
	"log"

	"hotel/api"
	"hotel/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

const (
	dburi = "mongodb://localhost:27017"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	// User initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
