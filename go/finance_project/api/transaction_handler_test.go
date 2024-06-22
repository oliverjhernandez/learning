package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"finance/fixtures"
	"finance/models"
)

func TestPostTx(t *testing.T) {
	app.Post("/tx", txHandler.HandlerPostTransaction)

	params := &models.CreateTransaction{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Relevance:   1,
		Currency:    models.COP,
		Account:     models.SAVINGS,
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
	app.Post("/tx", txHandler.HandlerPostTransaction)
	app.Get("/tx/:id", txHandler.HandlerGetTransaction)

	params := &models.CreateTransaction{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Relevance:   2,
		Currency:    models.COP,
		Account:     models.SAVINGS,
	}

	postTx, err := fixtures.AddTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	getTx, err := fixtures.GetTx(app, postTx.ID)
	if err != nil {
		t.Fatal(err)
	}

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
	app.Post("/tx", txHandler.HandlerPostTransaction)
	app.Delete("/tx/:id", txHandler.HandlerDeleteTransaction)
	app.Get("/tx/:id", txHandler.HandlerGetTransaction)

	params := &models.CreateTransaction{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Relevance:   3,
		Currency:    models.COP,
		Account:     models.SAVINGS,
	}

	postTx, err := fixtures.AddTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	delReq := httptest.NewRequest("DELETE", "/tx/"+string(rune(postTx.ID)), nil)
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

	var delTx models.Transaction
	json.Unmarshal(delRespJSON, &delTx)

	_, err = fixtures.GetTx(app, postTx.ID)

	var compareError ErrorMessage = NOT_FOUND
	if err.Error() != compareError.String() {
		t.Errorf("expected %s but got %s", compareError.String(), err.Error())
	}
}

func TestUpdateTx(t *testing.T) {
	app.Post("/tx/", txHandler.HandlerPostTransaction)
	app.Patch("/tx/:id", txHandler.HandlerUpdateTransaction)
	app.Get("/tx/:id", txHandler.HandlerGetTransaction)

	params := &models.CreateTransaction{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        1716308030,
		Relevance:   1,
		Currency:    models.COP,
		Account:     models.SAVINGS,
	}

	postTx, err := fixtures.AddTx(app, params)
	if err != nil {
		t.Fatal(err)
	}

	var accountType models.Account = models.CHECKINGS
	up := map[string]any{"account": accountType}
	b, _ := json.Marshal(up)

	upReq := httptest.NewRequest("PATCH", "/tx/"+string(rune(postTx.ID)), bytes.NewReader(b))
	upReq.Header.Add("Content-Type", "application/json")

	upResp, err := app.Test(upReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.ReadAll(upResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	getTx, err := fixtures.GetTx(app, postTx.ID)
	if err != nil {
		t.Fatal(err)
	}

	if getTx.Account != up["account"] {
		t.Errorf("expected %s but got %s", up["account"], getTx.Account)
	}
}
