package api

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"net/http/httptest"
// 	"testing"
//
// 	"finance/fixtures"
// 	"finance/models"
// )
//
// func TestPostTx(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
// 	app.Post("/account", accountHandler.HandlerPostAccount)
// 	app.Post("/tx", txHandler.HandlerPostTransaction)
//
// 	user, err := fixtures.AddUser(app)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	txn, err := fixtures.AddTx(app, user.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if txn.ID == 0 {
// 		t.Errorf("got %d but expected an actual id", txn.ID)
// 	}
// }
//
// func TestGetTx(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
// 	app.Post("/account", accountHandler.HandlerPostAccount)
// 	app.Post("/tx", txHandler.HandlerPostTransaction)
//
// 	user, err := fixtures.AddUser(app)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	acc, err := fixtures.AddAccount(app, user.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	txn, err := fixtures.GetTx(app, acc.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if txn.ID == 0 {
// 		t.Errorf("got %d but expected an actual id", txn.ID)
// 	}
// }
//
// func TestDeleteTx(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
// 	app.Post("/account", accountHandler.HandlerPostAccount)
// 	app.Post("/tx", txHandler.HandlerPostTransaction)
//
// 	user, err := fixtures.AddUser(app)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	acc, err := fixtures.AddAccount(app, user.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	newTxn, err := fixtures.AddTx(app, acc.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	delReq := httptest.NewRequest("DELETE", "/tx/"+string(rune(newTxn.ID)), nil)
// 	delReq.Header.Add("Content-Type", "application/json")
//
// 	delResp, err := app.Test(delReq, 1000*3)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var delRespJSON []byte
// 	delRespJSON, err = io.ReadAll(delResp.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var delTxn models.Transaction
// 	json.Unmarshal(delRespJSON, &delTxn)
//
// 	_, err = fixtures.GetTx(app, delTxn.ID)
//
// 	var compareError ErrorMessage = NOT_FOUND
// 	if err.Error() != compareError.String() {
// 		t.Errorf("expected %s but got %s", compareError.String(), err.Error())
// 	}
// }
//
// func TestUpdateTx(t *testing.T) {
// 	app.Post("/tx/", txHandler.HandlerPostTransaction)
// 	app.Post("/account/", txHandler.HandlerPostTransaction)
// 	app.Patch("/tx/:id", txHandler.HandlerUpdateTransaction)
// 	app.Get("/tx/:id", txHandler.HandlerGetTransaction)
//
// 	user, err := fixtures.AddUser(app)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	acc, err := fixtures.AddAccount(app, user.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	newTxn, err := fixtures.AddTx(app, acc.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	txn, err := fixtures.GetTx(app, newTxn.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var relevance models.Relevance = models.Important
// 	up := map[string]any{"account": relevance}
// 	b, _ := json.Marshal(up)
//
// 	upReq := httptest.NewRequest("PATCH", "/tx/"+string(rune(txn.ID)), bytes.NewReader(b))
// 	upReq.Header.Add("Content-Type", "application/json")
//
// 	upResp, err := app.Test(upReq, 1000*3)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	_, err = io.ReadAll(upResp.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	getTx, err := fixtures.GetTx(app, txn.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if getTx.Relevance.String() != up["relevance"] {
// 		t.Errorf("expected %s but got %s", up["account"], getTx.Relevance.String())
// 	}
// }
