package db

import (
	"context"

	"hotel/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	bookingCollection = "bookings"
)

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error)
	GetBookingByID(ctx context.Context, id string) (*types.Booking, error)
	UpdateBooking(ctx context.Context, id string, update bson.M) error
}

type MongoBookingStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoBookingStore(c *mongo.Client, dbname string) *MongoBookingStore {
	return &MongoBookingStore{
		client:     c,
		collection: c.Database(dbname).Collection(bookingCollection),
	}
}

func (bs *MongoBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	resp, err := bs.collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	booking.ID = resp.InsertedID.(primitive.ObjectID)

	return booking, nil
}

func (bs *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	cur, err := bs.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var bookings []*types.Booking
	if err := cur.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (bs *MongoBookingStore) GetBookingByID(ctx context.Context, id string) (*types.Booking, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking types.Booking
	if err := bs.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (bs *MongoBookingStore) UpdateBooking(ctx context.Context, id string, params bson.M) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": params,
	}
	_, err = bs.collection.UpdateByID(ctx, oid, update)
	return err
}
