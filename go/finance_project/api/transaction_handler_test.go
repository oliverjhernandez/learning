package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"finance/db"
	"finance/types"
)

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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.TDBURI))
	if err != nil {
		t.Fatal(err)
	}

	return &tDB{
		TransactionStore: db.NewMongoTransactionStore(client, db.TDBNAME),
	}
}

func insertTestTx(app *fiber.App, params *types.CreateTransactionParams) (*types.Transaction, error) {
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

	resp, err := insertTestTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Concept != params.Concept {
		t.Errorf("got %s but expected %s", resp.Concept, params.Concept)
	}

	if resp.Description != params.Description {
		t.Errorf("got %s but expected %s", resp.Description, params.Description)
	}

	if resp.Value != params.Value {
		t.Errorf("got %d but expected %d", resp.Value, params.Value)
	}

	if resp.Date != params.Date {
		t.Errorf("got %d but expected %d", resp.Date, params.Date)
	}

	if resp.Status != params.Status {
		t.Errorf("got %s but expected %s", resp.Status, params.Status)
	}

	if resp.Currency != params.Currency {
		t.Errorf("got %s but expected %s", resp.Currency, params.Currency)
	}

	if resp.Account != params.Account {
		t.Errorf("got %s but expected %s", resp.Account, params.Account)
	}
}

func TestGetTx(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	txHandler := NewTransactionHandler(tdb.TransactionStore)
	app := fiber.New()
	app.Post("/", txHandler.HandlerPostTransaction)
	app.Get("/:id", txHandler.HandlerGetTransaction)

	params := &types.CreateTransactionParams{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Status:      "Expendable",
		Currency:    "COP",
		Account:     "Savings",
	}

	postTx, err := insertTestTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/"+postTx.ID.Hex(), nil)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = app.Test(req, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	var respJSON []byte
	respJSON, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var getTx types.Transaction
	json.Unmarshal(respJSON, &getTx)

	if getTx.Concept != params.Concept {
		t.Errorf("got %s but expected %s", getTx.Concept, params.Concept)
	}

	if getTx.Description != params.Description {
		t.Errorf("got %s but expected %s", getTx.Description, params.Description)
	}

	if getTx.Value != params.Value {
		t.Errorf("got %d but expected %d", getTx.Value, params.Value)
	}

	if getTx.Date != params.Date {
		t.Errorf("got %d but expected %d", getTx.Date, params.Date)
	}

	if getTx.Status != params.Status {
		t.Errorf("got %s but expected %s", getTx.Status, params.Status)
	}

	if getTx.Currency != params.Currency {
		t.Errorf("got %s but expected %s", getTx.Currency, params.Currency)
	}

	if getTx.Account != params.Account {
		t.Errorf("got %s but expected %s", getTx.Account, params.Account)
	}
}

func TestDeleteTx(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	txHandler := NewTransactionHandler(tdb.TransactionStore)
	app := fiber.New()
	app.Post("/", txHandler.HandlerPostTransaction)
	app.Delete("/:id", txHandler.HandlerDeleteTransaction)
	app.Get("/:id", txHandler.HandlerGetTransaction)

	params := &types.CreateTransactionParams{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Status:      "Expendable",
		Currency:    "COP",
		Account:     "Savings",
	}

	postTx, err := insertTestTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	delReq := httptest.NewRequest("DELETE", "/"+postTx.ID.Hex(), nil)
	delReq.Header.Add("Content-Type", "application/json")

	delResp, err := app.Test(delReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	var delRespJSON []byte
	delRespJSON, err = io.ReadAll(delResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var delTx types.Transaction
	json.Unmarshal(delRespJSON, &delTx)

	getReq := httptest.NewRequest("GET", "/"+postTx.ID.Hex(), nil)
	getReq.Header.Add("Content-Type", "application/json")

	getResp, err := app.Test(getReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	var getRespJSON []byte
	getRespJSON, err = io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	const (
		MONGO_NO_DOCUMENTS = "mongo: no documents in result"
	)

	if string(getRespJSON) != MONGO_NO_DOCUMENTS {
		t.Errorf("expected %s but got %s", MONGO_NO_DOCUMENTS, string(getRespJSON))
	}
}

func TestUpdateTx(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	txHandler := NewTransactionHandler(tdb.TransactionStore)
	app := fiber.New()
	app.Post("/", txHandler.HandlerPostTransaction)
	app.Patch("/:id", txHandler.HandlerUpdateTransaction)
	app.Get("/:id", txHandler.HandlerGetTransaction)

	params := &types.CreateTransactionParams{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Status:      "Expendable",
		Currency:    "COP",
		Account:     "Savings",
	}

	postTx, err := insertTestTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	up := map[string]string{"account": "Checkings"}
	b, _ := json.Marshal(up)

	upReq := httptest.NewRequest("PATCH", "/"+postTx.ID.Hex(), bytes.NewReader(b))
	upReq.Header.Add("Content-Type", "application/json")

	upResp, err := app.Test(upReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.ReadAll(upResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	getReq := httptest.NewRequest("GET", "/"+postTx.ID.Hex(), nil)
	getReq.Header.Add("Content-Type", "application/json")

	getResp, err := app.Test(getReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	b, err = io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var getTx types.Transaction
	json.Unmarshal(b, &getTx)

	if getTx.Account != up["account"] {
		t.Errorf("expected %s but got %s", up["account"], getTx.Account)
	}
}
