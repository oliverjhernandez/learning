package db

import (
	"context"

	"finance/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection = "users"
)

type UserStore interface {
	GetUsers(ctx context.Context) ([]*types.User, error)
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
	UpdateUser(ctx context.Context, filter bson.M, params *types.UpdateUserParams) error
	DeleteUser(ctx context.Context, id string) error
}

type MongoUserStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func (us *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := us.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, users); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	var user types.User
	if err := us.collection.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := us.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (us *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params *types.UpdateUserParams) error {
	values := bson.D{
		{
			"$set", params.ToBSON(),
		},
	}
	_, err := us.collection.UpdateOne(ctx, filter, values)
	if err != nil {
		return err
	}

	return nil
}

func (us *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = us.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func NewMongoUserStore(mc *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client:     mc,
		dbname:     dbname,
		collection: mc.Database(dbname).Collection(userCollection),
	}
}
