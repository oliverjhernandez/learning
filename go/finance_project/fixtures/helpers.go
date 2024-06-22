package fixtures

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

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

func AddTx(app *fiber.App, params *models.CreateTransaction) (*models.Transaction, error) {
	tTx := models.NewTransactionFromParams(*params)

	b, _ := json.Marshal(tTx)

	req := httptest.NewRequest("POST", "/tx", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var tx models.Transaction
	json.NewDecoder(resp.Body).Decode(&tx)

	return &tx, nil
}

func GetTx(app *fiber.App, id int) (*models.Transaction, error) {
	req := httptest.NewRequest("GET", "/tx/"+string(rune(id)), nil)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var respJSON []byte
	respJSON, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getTx models.Transaction
	json.Unmarshal(respJSON, &getTx)

	return &getTx, nil
}

func AddUser(app *fiber.App, params *models.CreateUser) (*models.User, error) {
	tUser := models.NewUserFromParams(*params)

	b, _ := json.Marshal(tUser)
	req := httptest.NewRequest("POST", "/usr", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var user *models.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(app *fiber.App, id int) (*models.User, error) {
	req := httptest.NewRequest("GET", "/usr/"+string(rune(id)), nil)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err := app.Test(req, 1000*3)
	if err != nil {
		return nil, err
	}

	var respJSON []byte
	respJSON, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getUser models.User
	err = json.Unmarshal(respJSON, &getUser)
	if err != nil {
		return nil, err
	}

	return &getUser, nil
}
