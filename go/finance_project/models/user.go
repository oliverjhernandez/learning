package models

import (
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Passwd    string    `json:"-"`
	IsAdmin   bool      `json:"is_admin"`
	ID        int       `json:"id"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Passwd    string `json:"passwd"`
	IsAdmin   bool   `json:"is_admin"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPasswd(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}

func NewUserFromParams(u *CreateUser) (*User, error) {
	now := time.Now()
	encpw, err := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		CreatedAt: now,
		UpdatedAt: now,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		IsAdmin:   u.IsAdmin,
		Passwd:    string(encpw),
		Email:     u.Email,
	}, nil
}
