package fixtures

import (
	"context"
	"testing"

	"finance/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDbUri = "mongodb://localhost:27017"

type testDB struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testDB) TearDown(t *testing.T) {
	if err := tdb.client.Database(db.TDBNAME).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func Setup(t *testing.T) *testDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		t.Fatal(err)
	}

	userStore := db.NewMongoUserStore(client, db.TDBNAME)
	txStore := db.NewMongoTransactionStore(client, db.TDBNAME)

	return &testDB{
		client: client,
		Store: &db.Store{
			User: userStore,
			Tx:   txStore,
		},
	}
}
