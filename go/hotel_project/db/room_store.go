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
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (rs *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := rs.collection.InsertOne(ctx, room)
	if err != nil {
		log.Fatal(err)
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	// update hotel with this room id
	filter := bson.M{
		"_id": room.HotelID,
	}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}

	return room, nil
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(dbname).Collection(roomCollection),
	}
}
