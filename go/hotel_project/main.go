package main

import (
	"context"
	"errors"
	"flag"
	"log"

	api "hotel/api/handlers"
	"hotel/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new fiber instance with custom config
var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		// Send custom error page
		return ctx.JSON(map[string]string{"error": err.Error(), "code": string(code)})
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

	app.Get("/foo", handleFoo)

	// User initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "foo endpoint"})
}
