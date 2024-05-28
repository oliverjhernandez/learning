package main

import (
	"context"
	"fmt"
	"log"

	"hotel/db"
	"hotel/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(name, location string, rating int) error {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []types.Room{
		{
			Size:  "small",
			Price: 100.10,
		},
		{
			Size:  "normal",
			Price: 122.10,
		},
		{
			Size:  "king",
			Price: 269.10,
		},
	}
	newHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range rooms {
		v.HotelID = newHotel.ID
		newRoom, err := roomStore.InsertRoom(ctx, &v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", newRoom)
	}
	return nil
}

func main() {
	seedHotel("Bellucia", "France", 5)
	seedHotel("Beetlejuice", "The Netherlands", 4)
	seedHotel("Stabby Stabby", "England", 2)
}

func init() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client, db.DBNAME)
	roomStore = db.NewMongoRoomStore(client, hotelStore, db.DBNAME)
}
