package fixtures

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"

	"finance/db"
	"finance/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDbUri = "mongodb://localhost:27017"

type TestStore struct {
	client *mongo.Client
	*db.Store
}

func (ts *TestStore) TearDown() error {
	if err := ts.client.Database(db.TDBNAME).Drop(context.TODO()); err != nil {
		return err
	}

	return nil
}

func NewTestStore() (*TestStore, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		return nil, err
	}

	userStore := db.NewMongoUserStore(client, db.TDBNAME)
	txStore := db.NewMongoTransactionStore(client, db.TDBNAME)
	creditStore := db.NewMongoCreditStore(client, db.TDBNAME)

	return &TestStore{
		client: client,
		Store: &db.Store{
			User:   userStore,
			Tx:     txStore,
			Credit: creditStore,
		},
	}, nil
}

func AddTx(app *fiber.App, params *types.CreateTransactionParams) (*types.Transaction, error) {
	tTx, err := types.NewTransactionFromParams(*params)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(tTx)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var tx types.Transaction
	json.NewDecoder(resp.Body).Decode(&tx)

	return &tx, nil
}

func AddUser(app *fiber.App, params *types.CreateUserParams) (*types.User, error) {
	tUser, err := types.NewUserFromParams(*params)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(tUser)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		return nil, err
	}

	var user *types.User
	json.NewDecoder(resp.Body).Decode(&user)

	return user, nil
}
