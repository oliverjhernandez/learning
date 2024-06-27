package fixtures

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"finance/db"
	"finance/models"

	"github.com/gofiber/fiber/v2"
)

type TestStore struct {
	Store *db.Store
}

func (ts *TestStore) TearDown() error {
	if err := ts.Store.Close(); err != nil {
		return err
	}

	return nil
}

func NewTestStore() (*TestStore, error) {
	store, _, err := db.NewStore()
	if err != nil {
		return nil, err
	}

	return &TestStore{
		Store: store,
	}, nil
}

func AddUser(app *fiber.App) (*models.User, error) {
	params := models.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corinapulido@gmail.com",
	}

	newUser, err := models.NewUserFromParams(&params)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(newUser)
	req := httptest.NewRequest("POST", "/usr", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(app *fiber.App, userID int) (*models.User, error) {
	req := httptest.NewRequest("GET", "/usr/"+string(rune(userID)), nil)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func AddTx(app *fiber.App, accID int) (*models.Transaction, error) {
	params := models.CreateTransaction{
		Concept:     "Alquiler",
		Description: "Castel D'Aiano",
		Value:       2250000,
		Date:        time.Now().Add(time.Hour * 24 * -7),
		Relevance:   1,
		AccountID:   accID,
	}

	newTxn, err := models.NewTransactionFromParams(params)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(newTxn)

	req := httptest.NewRequest("POST", "/tx", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var txn models.Transaction
	if err := json.NewDecoder(resp.Body).Decode(&txn); err != nil {
		return nil, err
	}

	return &txn, nil
}

func GetTx(app *fiber.App, txnID int) (*models.Transaction, error) {
	req := httptest.NewRequest("GET", "/tx/"+string(rune(txnID)), nil)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var txn models.Transaction
	if err := json.NewDecoder(resp.Body).Decode(&txn); err != nil {
		return nil, err
	}

	return &txn, nil
}

func AddAccount(app *fiber.App, userID int) (*models.Account, error) {
	params := models.CreateAccount{
		Name:     "Main",
		Entity:   models.BANCOLOMBIA,
		Currency: models.COP,
		UserID:   userID,
	}

	acc, err := models.NewAccountFromParams(&params)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(acc)
	req := httptest.NewRequest("POST", "/account", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var account models.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func GetAccount(app *fiber.App, accID int) (*models.Account, error) {
	req := httptest.NewRequest("GET", "/account/"+string(rune(accID)), nil)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var account models.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return nil, err
	}

	return &account, nil
}
