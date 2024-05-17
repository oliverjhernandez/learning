package db

import (
	"context"

	"finance/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TransactionCollection = "transaction"
)

type TransactionStore interface {
	GetTransactionByID(ctx context.Context, id string) (*types.Transaction, error)
	InsertTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error)
}

type MongoTransactionStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func (ts *MongoTransactionStore) GetTransactionByID(ctx context.Context, id string) (*types.Transaction, error) {
	var c context.Context

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var tx types.Transaction
	if err := ts.collection.FindOne(c, bson.M{"_id": oid}).Decode(tx); err != nil {
		return nil, err
	}

	return &tx, nil
}

func (ts *MongoTransactionStore) InsertTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	res, err := ts.collection.InsertOne(ctx, tx)
	if err != nil {
		return nil, err
	}
	tx.ID = res.InsertedID.(primitive.ObjectID)
	return tx, nil
}

func NewMongoTransactionStore(c *mongo.Client) *MongoTransactionStore {
	return &MongoTransactionStore{
		client:     c,
		dbname:     DBNAME,
		collection: c.Database(DBNAME).Collection(TransactionCollection),
	}
}
