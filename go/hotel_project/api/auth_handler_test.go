package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hotel/db/fixtures"

	"github.com/gofiber/fiber/v2"
)

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	newUser := fixtures.AddUser(tdb.Store, "James", "Foo", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:  newUser.Email,
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
	newUser := fixtures.AddUser(tdb.Store, "James", "Foo", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:  newUser.Email,
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
