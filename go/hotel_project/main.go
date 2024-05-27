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

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	app := fiber.New(config)
	appv1 := app.Group("/api/v1")

	// User initialization
	var (
		userStore   = db.NewMongoUserStore(client, db.DBNAME)
		userHandler = api.NewUserHandler(userStore)
	)

	appv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	appv1.Put("/user/:id", userHandler.HandleUpdateUser)
	appv1.Get("/user", userHandler.HandleGetUsers)
	appv1.Post("/user", userHandler.HandlePostUser)
	appv1.Get("/user/:id", userHandler.HandleGetUser)

	// Hotels
	var (
		hotelStore   = db.NewMongoHotelStore(client, db.DBNAME)
		roomStore    = db.NewMongoRoomStore(client, hotelStore, db.DBNAME)
		hotelHandler = api.NewHotelHandler(hotelStore, roomStore)
	)
	appv1.Get("/hotels", hotelHandler.HandlerGetHotels)

	app.Listen(*listenAddr)
}
