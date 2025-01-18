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
// func TestPostUser(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
//
// 	user, err := fixtures.AddUser(app)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if user.ID == 0 {
// 		t.Errorf("got %d but expected an actual id", user.ID)
// 	}
// }
//
// func TestGetUser(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
// 	app.Get("/usr/:id", userHandler.HandlerGetUser)
//
// 	newUser, err := fixtures.AddUser(app)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	user, err := fixtures.GetUser(app, newUser.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if user.ID == 0 {
// 		t.Errorf("got %d but expected an actual id", user.ID)
// 	}
// }
//
// func TestUpdateUser(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
// 	app.Patch("/usr/:id", userHandler.HandlerUpdateUser)
// 	app.Get("/usr/:id", userHandler.HandlerGetUser)
//
// 	user, err := fixtures.GetUser(app, newUser.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	name := "Mi Vida"
// 	update := models.UpdateUser{FirstName: name}
//
// 	b, err := json.Marshal(update)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	upReq := httptest.NewRequest("PATCH", "/usr/"+string(rune(postUser.ID)), bytes.NewReader(b))
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
// 	getUser, err := fixtures.GetUser(app, postUser.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if getUser.FirstName != name {
// 		t.Errorf("expected %s but got %s", name, getUser.FirstName)
// 	}
// }
//
// func TestDeleteUser(t *testing.T) {
// 	app.Post("/usr", userHandler.HandlerPostUser)
// 	app.Delete("/usr/:id", userHandler.HandlerDeleteUser)
// 	app.Get("/usr/:id", userHandler.HandlerGetUser)
//
// 	params := &models.CreateUser{
// 		FirstName: "Corina",
// 		LastName:  "Pulido",
// 		Email:     "corinapulido@gmail.com",
// 	}
//
// 	postUser, err := fixtures.AddUser(app, params)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	delReq := httptest.NewRequest("DELETE", "/usr/"+string(rune(postUser.ID)), nil)
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
// 	var delUser models.User
// 	json.Unmarshal(delRespJSON, &delUser)
//
// 	_, err = fixtures.GetUser(app, postUser.ID)
//
// 	var not_found ErrorMessage = NOT_FOUND
// 	if err.Error() != not_found.String() {
// 		t.Errorf("expected %s but got %s", not_found.String(), err.Error())
// 	}
// }
