package main

import (
	"context"
	"fmt"
	"log"

	"hotel/db"
	"hotel/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Bellucia",
		Location: "Narnia",
	}

	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 100.10,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 150.10,
		},
		{
			Type:      types.DeluxRoomType,
			BasePrice: 300.10,
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

	fmt.Printf("%+v", newHotel)
}
