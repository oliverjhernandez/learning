package api

import (
	"context"
	"testing"

	"hotel/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDbUri = "mongodb://localhost:27017"

type testDB struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testDB) tearDown(t *testing.T) {
	if err := tdb.client.Database(db.TDBNAME).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		t.Fatal(err)
	}

	userStore := db.NewMongoUserStore(client, db.TDBNAME)
	hotelStore := db.NewMongoHotelStore(client, db.TDBNAME)
	roomStore := db.NewMongoRoomStore(client, hotelStore, db.TDBNAME)
	bookingStore := db.NewMongoBookingStore(client, db.TDBNAME)

	return &testDB{
		client: client,
		Store: &db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		},
	}
}
