package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"finance/db"
	"finance/types"
)

const tDBURI = "mongodb://localhost:27017"

type tDB struct {
	db.TransactionStore
}

func (tdb *tDB) tearDown(t *testing.T) {
	fmt.Println("--- dropping database")
	if err := tdb.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *tDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(tDBURI))
	if err != nil {
		t.Fatal(err)
	}

	return &tDB{
		TransactionStore: db.NewMongoTransactionStore(client, db.TDBNAME),
	}
}

func TestPostTx(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	txHandler := NewTransactionHandler(tdb.TransactionStore)

	app.Post("/", txHandler.HandlerPostTransaction)

	params := &types.CreateTransactionParams{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Status:      "Expendable",
		Currency:    "COP",
		Account:     "Savings",
	}

	tTx, err := types.NewTransactionFromParams(*params)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(tTx)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	var tx types.Transaction
	json.NewDecoder(resp.Body).Decode(&tx)

	if tx.Concept != params.Concept {
		t.Errorf("got %s but expected %s", tx.Concept, params.Concept)
	}

	if tx.Description != params.Description {
		t.Errorf("got %s but expected %s", tx.Description, params.Description)
	}

	if tx.Value != params.Value {
		t.Errorf("got %d but expected %d", tx.Value, params.Value)
	}

	if tx.Date != params.Date {
		t.Errorf("got %d but expected %d", tx.Date, params.Date)
	}

	if tx.Status != params.Status {
		t.Errorf("got %s but expected %s", tx.Status, params.Status)
	}

	if tx.Currency != params.Currency {
		t.Errorf("got %s but expected %s", tx.Currency, params.Currency)
	}

	if tx.Account != params.Account {
		t.Errorf("got %s but expected %s", tx.Account, params.Account)
	}
}
