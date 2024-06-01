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

		userStore    = db.NewMongoUserStore(client, db.DBNAME)
		hotelStore   = db.NewMongoHotelStore(client, db.DBNAME)
		bookingStore = db.NewMongoBookingStore(client, db.DBNAME)
		roomStore    = db.NewMongoRoomStore(client, hotelStore, db.DBNAME)
		store        = &db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewBookRoomHandler(store)
		hotelHandler   = api.NewHotelHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		userHandler    = api.NewUserHandler(store)

		appv1 = app.Group("/v1", middleware.JWTAuthentication(store.User))
		auth  = app.Group("/api")
		admin = appv1.Group("/admin", middleware.AdminAuth)
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

	// Rooms
	appv1.Get("/room", roomHandler.HandleGetRooms)
	appv1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	// Bookings
	// TODO: cancel booking
	admin.Get("/booking", bookingHandler.HandlerGetBookings)
	appv1.Get("/booking/:id", bookingHandler.HandlerGetBooking)

	app.Listen(*listenAddr)
}
