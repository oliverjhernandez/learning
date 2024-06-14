package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"finance/fixtures"
	"finance/types"
)

func TestPostTx(t *testing.T) {
	app.Post("/", txHandler.HandlerPostTransaction)

	params := &types.CreateTransactionParams{
		TransactionBase: types.TransactionBase{
			Concept:     "Alquiler",
			Description: "Castel D'Aiano",
			Value:       2250000,
			Date:        1716308030,
			Relevance:   1,
			Currency:    types.COP,
			Account:     types.SAVINGS,
		},
	}

	resp, err := fixtures.AddTx(app, params)
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

	if resp.Relevance != params.Relevance {
		t.Errorf("got %s but expected %s", resp.Relevance, params.Relevance)
	}

	if resp.Currency != params.Currency {
		t.Errorf("got %s but expected %s", resp.Currency, params.Currency)
	}

	if resp.Account != params.Account {
		t.Errorf("got %s but expected %s", resp.Account, params.Account)
	}
}

func TestGetTx(t *testing.T) {
	app.Post("/", txHandler.HandlerPostTransaction)
	app.Get("/:id", txHandler.HandlerGetTransaction)

	params := &types.CreateTransactionParams{
		TransactionBase: types.TransactionBase{
			Concept:     "Alquiler",
			Description: "Castel D'Aiano",
			Value:       2250000,
			Date:        1716308030,
			Relevance:   2,
			Currency:    types.COP,
			Account:     types.SAVINGS,
		},
	}

	postTx, err := fixtures.AddTx(app, params)
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

	if getTx.Relevance != params.Relevance {
		t.Errorf("got %s but expected %s", getTx.Relevance, params.Relevance)
	}

	if getTx.Currency != params.Currency {
		t.Errorf("got %s but expected %s", getTx.Currency, params.Currency)
	}

	if getTx.Account != params.Account {
		t.Errorf("got %s but expected %s", getTx.Account, params.Account)
	}
}

func TestDeleteTx(t *testing.T) {
	app.Post("/", txHandler.HandlerPostTransaction)
	app.Delete("/:id", txHandler.HandlerDeleteTransaction)
	app.Get("/:id", txHandler.HandlerGetTransaction)

	params := &types.CreateTransactionParams{
		TransactionBase: types.TransactionBase{
			Concept:     "Alquiler",
			Description: "Castel D'Aiano",
			Value:       2250000,
			Date:        1716308030,
			Relevance:   3,
			Currency:    types.COP,
			Account:     types.SAVINGS,
		},
	}

	postTx, err := fixtures.AddTx(app, params)
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

	var compareError ErrorMessage = NOT_FOUND
	if string(getRespJSON) != compareError.String() {
		t.Errorf("expected %s but got %s", compareError.String(), string(getRespJSON))
	}
}

func TestUpdateTx(t *testing.T) {
	app.Post("/", txHandler.HandlerPostTransaction)
	app.Patch("/:id", txHandler.HandlerUpdateTransaction)
	app.Get("/:id", txHandler.HandlerGetTransaction)

	params := &types.CreateTransactionParams{
		TransactionBase: types.TransactionBase{
			Concept:     "Alquiler",
			Description: "Castel D'Aiano",
			Value:       2250000,
			Date:        1716308030,
			Relevance:   1,
			Currency:    types.COP,
			Account:     types.SAVINGS,
		},
	}

	postTx, err := fixtures.AddTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	var accountType types.Account = types.SAVINGS
	up := map[string]any{"account": accountType}
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
