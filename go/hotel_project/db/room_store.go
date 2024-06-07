package db

import (
	"context"
	"log"

	"hotel/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	roomCollection = "rooms"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error)
}

type MongoRoomStore struct {
	HotelStore
	client     *mongo.Client
	collection *mongo.Collection
}

func (rs *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := rs.collection.InsertOne(ctx, room)
	if err != nil {
		log.Fatal(err)
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)
	// update hotel with his room id
	filter := Params{"_id": room.HotelID}
	update := Params{"$push": bson.M{"rooms": room.ID}}
	if err := rs.UpdateHotel(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}

func (rs *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	resp, err := rs.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var rooms []*types.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func NewMongoRoomStore(c *mongo.Client, hs HotelStore, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client:     c,
		collection: c.Database(dbname).Collection(roomCollection),
		HotelStore: hs,
	}
}
