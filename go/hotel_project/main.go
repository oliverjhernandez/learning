package main

import (
	"context"
	"flag"
	"log"

	"hotel/api"
	"hotel/api/middleware"
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

	var (
		listenAddr = flag.String("listenAddr", ":3000", "The listen address of the API server")
		app        = fiber.New(config)
		auth       = app.Group("/api")
		appv1      = app.Group("/v1", middleware.JWTAuthentication)

		userStore  = db.NewMongoUserStore(client, db.DBNAME)
		hotelStore = db.NewMongoHotelStore(client, db.DBNAME)
		roomStore  = db.NewMongoRoomStore(client, hotelStore, db.DBNAME)
		store      = &db.Store{
			User:  userStore,
			Hotel: hotelStore,
			Room:  roomStore,
		}
		authHandler  = api.NewAuthHandler(userStore)
		hotelHandler = api.NewHotelHandler(store)
		userHandler  = api.NewUserHandler(store)
	)

	// Auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// Versioned API routes
	// Users
	appv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	appv1.Put("/user/:id", userHandler.HandleUpdateUser)
	appv1.Get("/user", userHandler.HandleGetUsers)
	appv1.Post("/user", userHandler.HandlePostUser)
	appv1.Get("/user/:id", userHandler.HandleGetUser)

	// Hotels
	appv1.Get("/hotel", hotelHandler.HandlerGetHotels)
	appv1.Get("/hotel/:id", hotelHandler.HandlerGetHotel)
	appv1.Get("/hotels/:id/rooms", hotelHandler.HandlerGetRooms)

	app.Listen(*listenAddr)
}
