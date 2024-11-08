package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"money_lovers/internal/db"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(store *db.Store) *AuthHandler {
	return &AuthHandler{
		userStore: store.UserStore,
	}
}

type AuthParams struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

type AuthResponse struct {
	User  *db.User `json:"user"`
	Token string   `json:"token"`
}

type genericResponse struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func (ah *AuthHandler) HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var params AuthParams
	if err := readJSON(r, &params); err != nil {
		unauthorizedError(err)
		return
	}

	user, err := ah.userStore.GetUserByEmail(c, nil, params.Email)
	if err != nil {
		unauthorizedError(err)
		return
	}

	// TODO: Should be better to send this in HTTP headers
	resp := AuthResponse{
		User:  user,
		Token: CreateTokenFromUser(user),
	}

	fmt.Println("Authenticated -> ", user.FirstName)

	err = writeJSON(w, resp)
	if err != nil {
		internalServerError(err)
		return
	}

	return
}

func CreateTokenFromUser(user *db.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4).Unix()
	// TODO: Apparently there are standard claim names
	// Let's investigate that and use it
	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret", err)
	}

	return tokenStr
}
