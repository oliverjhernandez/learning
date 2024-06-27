package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"finance/db"
	"finance/models"

	"github.com/gofiber/fiber/v2"
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
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

type genericResponse struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func invalidCredentials(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(genericResponse{
		Type: "error",
		Msg:  "invalid credentials",
	})
}

func (ah *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := ah.userStore.GetUserByEmail(c.Context(), nil, params.Email)
	if err != nil {
		return err
	}

	if !models.IsValidPasswd(user.Passwd, params.Passwd) {
		return invalidCredentials(c)
	}

	// TODO: Should be better to send this in HTTP headers
	resp := AuthResponse{
		User:  &user,
		Token: CreateTokenFromUser(&user),
	}

	fmt.Println("Authenticated -> ", user.FirstName)
	return c.JSON(resp)
}

func CreateTokenFromUser(user *models.User) string {
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
