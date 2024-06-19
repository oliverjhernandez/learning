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

func TestPostUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)

	params := &types.CreateUserParams{
		UserBase: types.UserBase{
			FirstName: "Corina",
			LastName:  "Pulido",
			Email:     "corinapulido@gmail.com",
		},
	}

	resp, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	if resp.UserBase != params.UserBase {
		t.Errorf("got %v but expected %v", resp.UserBase, params.UserBase)
	}
}

func TestGetUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)
	app.Get("/usr/:id", userHandler.HandlerGetUser)

	params := &types.CreateUserParams{
		UserBase: types.UserBase{
			FirstName: "Corina",
			LastName:  "Pulido",
			Email:     "corinapulido@gmail.com",
		},
	}

	postUser, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/usr/"+postUser.ID.Hex(), nil)
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

	var getUser types.User
	err = json.Unmarshal(respJSON, &getUser)
	if err != nil {
		t.Fatal(err)
	}

	if getUser.UserBase != params.UserBase {
		t.Errorf("got %s but expected %s", getUser.UserBase, params.UserBase)
	}
}

func TestUpdateUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)
	app.Patch("/usr/:id", userHandler.HandlerUpdateUser)
	app.Get("/usr/:id", userHandler.HandlerGetUser)

	params := &types.CreateUserParams{
		UserBase: types.UserBase{
			FirstName: "Corina",
			LastName:  "Pulido",
			Email:     "corinapulido@gmail.com",
		},
	}

	postUser, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	name := "Mi Vida"
	up := types.UserBase{FirstName: name}
	updateParams := types.UpdateUserParams{
		UserBase: up,
	}

	b, err := json.Marshal(updateParams)
	if err != nil {
		t.Fatal(err)
	}

	upReq := httptest.NewRequest("PATCH", "/usr/"+postUser.ID.Hex(), bytes.NewReader(b))
	upReq.Header.Add("Content-Type", "application/json")

	upResp, err := app.Test(upReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.ReadAll(upResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	getReq := httptest.NewRequest("GET", "/usr/"+postUser.ID.Hex(), nil)
	getReq.Header.Add("Content-Type", "application/json")

	getResp, err := app.Test(getReq, 1000*3)
	if err != nil {
		t.Fatal(err)
	}

	b, err = io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var user types.User
	json.Unmarshal(b, &user)

	if user.UserBase.FirstName != name {
		t.Errorf("expected %s but got %s", name, user.UserBase.FirstName)
	}
}

func TestDeleteUser(t *testing.T) {
	app.Post("/usr", userHandler.HandlerPostUser)
	app.Delete("/usr/:id", userHandler.HandlerDeleteUser)
	app.Get("/usr/:id", userHandler.HandlerGetUser)

	params := &types.CreateUserParams{
		UserBase: types.UserBase{
			FirstName: "Corina",
			LastName:  "Pulido",
			Email:     "corinapulido@gmail.com",
		},
	}

	resp, err := fixtures.AddUser(app, params)
	if err != nil {
		t.Fatal(err)
	}

	delReq := httptest.NewRequest("DELETE", "/usr/"+resp.ID.Hex(), nil)
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

	var delUser types.User
	json.Unmarshal(delRespJSON, &delUser)

	getReq := httptest.NewRequest("GET", "/usr/"+resp.ID.Hex(), nil)
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
