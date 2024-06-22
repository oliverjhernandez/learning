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

func TestPostUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)

	params := &models.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corinapulido@gmail.com",
	}

	resp, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	if resp.FirstName != params.FirstName {
		t.Errorf("got %v but expected %v", resp.FirstName, params.FirstName)
	}

	if resp.LastName != params.LastName {
		t.Errorf("got %v but expected %v", resp.LastName, params.LastName)
	}

	if resp.Email != params.Email {
		t.Errorf("got %v but expected %v", resp.Email, params.Email)
	}
}

func TestGetUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)
	app.Get("/usr/:id", userHandler.HandlerGetUser)

	params := &models.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corinapulido@gmail.com",
	}

	postUser, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	getUser, err := fixtures.GetUser(app, postUser.ID)
	if err != nil {
		t.Fatal(err)
	}

	if getUser.FirstName != params.FirstName {
		t.Errorf("got %s but expected %s", getUser.FirstName, params.FirstName)
	}

	if getUser.LastName != params.LastName {
		t.Errorf("got %s but expected %s", getUser.LastName, params.LastName)
	}

	if getUser.Email != params.Email {
		t.Errorf("got %s but expected %s", getUser.Email, params.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)
	app.Patch("/usr/:id", userHandler.HandlerUpdateUser)
	app.Get("/usr/:id", userHandler.HandlerGetUser)

	params := &models.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corinapulido@gmail.com",
	}

	postUser, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	name := "Mi Vida"
	update := models.UpdateUser{FirstName: name}

	b, err := json.Marshal(update)
	if err != nil {
		t.Fatal(err)
	}

	upReq := httptest.NewRequest("PATCH", "/usr/"+string(rune(postUser.ID)), bytes.NewReader(b))
	upReq.Header.Add("Content-Type", "application/json")

	upResp, err := app.Test(upReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.ReadAll(upResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	getUser, err := fixtures.GetUser(app, postUser.ID)
	if err != nil {
		t.Fatal(err)
	}

	if getUser.FirstName != name {
		t.Errorf("expected %s but got %s", name, getUser.FirstName)
	}
}

func TestDeleteUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)
	app.Delete("/usr/:id", userHandler.HandlerDeleteUser)
	app.Get("/usr/:id", userHandler.HandlerGetUser)

	params := &models.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corinapulido@gmail.com",
	}

	postUser, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	delReq := httptest.NewRequest("DELETE", "/usr/"+string(rune(postUser.ID)), nil)
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

	var delUser models.User
	json.Unmarshal(delRespJSON, &delUser)

	_, err = fixtures.GetUser(app, postUser.ID)

	var not_found ErrorMessage = NOT_FOUND
	if err.Error() != not_found.String() {
		t.Errorf("expected %s but got %s", not_found.String(), err.Error())
	}
}
