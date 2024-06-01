package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hotel/db"
	"hotel/types"

	"github.com/gofiber/fiber/v2"
)

func insertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	params := types.CreateUserParams{
		FirstName: "James",
		LastName:  "Hetfield",
		Passwd:    "superarrecho",
		Email:     "metallica.nice.com",
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := userStore.InsertUser(context.TODO(), user); err != nil {
		t.Fatal(err)
	}

	return user
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	insertTestUser(t, tdb.User)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:  "metallica.nice.com",
		Passwd: "superarrecho",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected http status 200 but got %d", resp.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Error(err)
	}

	if authResp.Token == "" {
		t.Fatal("expected the jwt token to be present in the auth response")
	}
}

func TestAuthenticateFailure(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	insertTestUser(t, tdb.User)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:  "metallica.nice.com",
		Passwd: "supervergatario",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected http status 200 but got %d", resp.StatusCode)
	}

	var genResp genericResponse
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		t.Error(err)
	}

	if genResp.Type != "error" {
		t.Fatalf("expected type to be <error> but got %s", genResp.Type)
	}

	if genResp.Msg != "invalid credentials" {
		t.Fatalf("expected msg to be <invalid credentials> but got %s", genResp.Msg)
	}
}
