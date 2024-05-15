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
}

type MongoTransactionStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func (ts *MongoTransactionStore) GetTransactionByID(ctx context.Context, id string) (*types.Transaction, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var tx types.Transaction
	if err := ts.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(tx); err != nil {
		return nil, err
	}

	return &tx, nil
}

func NewMongoTransactionStore(c *mongo.Client) *MongoTransactionStore {
	return &MongoTransactionStore{
		client:     c,
		dbname:     DBNAME,
		collection: c.Database(DBNAME).Collection(TransactionCollection),
	}
}
