package db

import (
	"context"

	"finance/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	transactionCollection = "transactions"
)

type Dropper interface {
	Drop(ctx context.Context) error
}

type TransactionStore interface {
	Dropper

	GetTransactions(ctx context.Context) ([]*types.Transaction, error)
	GetTransactionByID(ctx context.Context, id string) (*types.Transaction, error)
	InsertTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error)
	UpdateTransaction(ctx context.Context, filter bson.M, params *types.UpdateTransactionParams) error
	DeleteTransaction(ctx context.Context, id string) error
}

type MongoTransactionStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	dbname     string
}

func (ts *MongoTransactionStore) Drop(ctx context.Context) error {
	return ts.collection.Drop(ctx)
}

func (ts *MongoTransactionStore) GetTransactions(ctx context.Context) ([]*types.Transaction, error) {
	cur, err := ts.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var txs []*types.Transaction
	if err := cur.All(ctx, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}

func (ts *MongoTransactionStore) GetTransactionByID(ctx context.Context, id string) (*types.Transaction, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var tx types.Transaction
	if err := ts.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&tx); err != nil {
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

func (ts *MongoTransactionStore) DeleteTransaction(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = ts.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (ts *MongoTransactionStore) UpdateTransaction(ctx context.Context, filter bson.M, params *types.UpdateTransactionParams) error {
	values := bson.D{
		primitive.E{
			Key: "$set", Value: params.ToBSON(),
		},
	}

	_, err := ts.collection.UpdateOne(ctx, filter, values)
	if err != nil {
		return err
	}

	return nil
}

func NewMongoTransactionStore(mc *mongo.Client, dbname string) *MongoTransactionStore {
	return &MongoTransactionStore{
		client:     mc,
		dbname:     dbname,
		collection: mc.Database(dbname).Collection(transactionCollection),
	}
}
