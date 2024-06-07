package db

import (
	"context"

	"hotel/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	hotelCollection = "hotels"
)

type HotelStore interface {
	GetHotels(ctx context.Context, filter Params) ([]*types.Hotel, error)
	GetHotelByID(ctx context.Context, id string) (*types.Hotel, error)
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, filter Params, values Params) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (hs *MongoHotelStore) GetHotelByID(ctx context.Context, id string) (*types.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := Params{"_id": oid}
	var hotel types.Hotel
	if err := hs.collection.FindOne(ctx, filter).Decode(&hotel); err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (hs *MongoHotelStore) GetHotels(ctx context.Context, filter Params) ([]*types.Hotel, error) {
	resp, err := hs.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (hs *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := hs.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (hs *MongoHotelStore) UpdateHotel(ctx context.Context, filter Params, params Params) error {
	_, err := hs.collection.UpdateOne(ctx, filter, params)
	return err
}

func NewMongoHotelStore(c *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client:     c,
		collection: c.Database(dbname).Collection(hotelCollection),
	}
}
