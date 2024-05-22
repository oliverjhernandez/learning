package db

import (
	"context"

	"hotel/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	hotelCollection = "hotels"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, filter bson.M, values bson.M) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (hs *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := hs.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (hs *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, params bson.M) error {
	_, err := hs.collection.UpdateOne(ctx, filter, params)
	return err
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(dbname).Collection(hotelCollection),
	}
}
