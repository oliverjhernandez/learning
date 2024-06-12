package db

import (
	"context"

	"finance/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	creditsCollection = "credits"
)

type MongoCreditStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	dbname     string
}

func NewMongoCreditStore(mc *mongo.Client, dbname string) *MongoCreditStore {
	return &MongoCreditStore{
		client:     mc,
		collection: mc.Database(dbname).Collection(creditsCollection),
		dbname:     dbname,
	}
}

func (cs *MongoCreditStore) GetCredits(ctx context.Context) ([]*types.Credit, error) {
	cur, err := cs.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var credits []*types.Credit
	if err := cur.All(ctx, &credits); err != nil {
		return nil, err
	}
	return credits, nil
}

func (cs *MongoCreditStore) GetCreditByID(ctx context.Context, id string) (*types.Credit, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": oid,
	}

	var credit types.Credit
	if err := cs.collection.FindOne(ctx, filter).Decode(&credit); err != nil {
		return nil, err
	}

	return &credit, nil
}

func (cs *MongoCreditStore) InsertCredit(ctx context.Context, cred *types.Credit) (*types.Credit, error) {
	res, err := cs.collection.InsertOne(ctx, cred)
	if err != nil {
		return nil, err
	}
	cred.ID = res.InsertedID.(primitive.ObjectID)

	return cred, nil
}

func (cs *MongoCreditStore) UpdateCredit(ctx context.Context, filter Params, update *types.UpdateCreditParams) error {
	updateQuery := bson.D{
		primitive.E{
			Key: "$set", Value: update,
		},
	}
	_, err := cs.collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err
	}

	return nil
}

func (cs *MongoCreditStore) DeleteCreditByID(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}
	_, err = cs.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
