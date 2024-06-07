package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"hotel/api"
	"hotel/db"
	"hotel/db/fixtures"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, hotelStore, db.DBNAME)
	userStore := db.NewMongoUserStore(client, db.DBNAME)
	bookingStore := db.NewMongoBookingStore(client, db.DBNAME)

	store := db.Store{
		User:    userStore,
		Hotel:   hotelStore,
		Room:    roomStore,
		Booking: bookingStore,
	}

	user := fixtures.AddUser(&store, "Oliver", "Hernandez", false)
	fmt.Printf("%s -> %s\n", user.FirstName, api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(&store, "Admin", "Admin", true)
	fmt.Printf("%s -> %s\n", admin.FirstName, api.CreateTokenFromUser(user))
	_ = admin

	hotel := fixtures.AddHotel(&store, "stabby stabby", "bermuda", 5, nil)
	room := fixtures.AddRoom(&store, "large", true, 189.99, hotel.ID)
	booking := fixtures.AddBooking(
		&store,
		user.ID,
		room.ID,
		time.Date(2024, time.June, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2024, time.July, 5, 0, 0, 0, 0, time.UTC),
	)
	fmt.Println("booking -> ", booking.ID)
}
