package db

import (
	"errors"
	"time"

	"casita/internal/validator"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Passwd    password  `json:"-"`
	Version   string    `json:"-"`
	Activated bool      `json:"-"`
}

type CreateUser struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Passwd    password `json:"passwd"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintext string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintext
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintext))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err

		}
	}

	return true, nil
}

///////////////////////////

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	// v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, plaintext string) {
	v.Check(plaintext != "", "passwd", "must be provided")
	v.Check(len(plaintext) >= 8, "passwd", "must be at least 8 bytes long")
	v.Check(len(plaintext) <= 72, "passwd", "must not be longer than 72 bytes")
}

func NewUserFromParams(u *CreateUser) (*User, error) {
	now := time.Now()
	encpw, err := bcrypt.GenerateFromPassword([]byte(u.Passwd.hash), bcryptCost)
	if err != nil {
		return nil, err
	}

	passwd := &password{
		plaintext: u.Passwd.plaintext,
		hash:      encpw,
	}

	return &User{
		CreatedAt: now,
		UpdatedAt: now,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Passwd:    *passwd,
	}, nil
}

func ValidateUser(v *validator.Validator, u *User) {
	// FirstName
	v.Check(u.FirstName != "", "first_name", "must be provided")
	v.Check(len(u.FirstName) <= 500, "first_name", "must be at least 3 bytes long")

	// LastName
	v.Check(u.LastName != "", "last_name", "must be provided")
	v.Check(len(u.LastName) >= 3, "last_name", "must be at least 3 bytes long")

	// Validate email
	ValidateEmail(v, u.Email)

	if u.Passwd.plaintext != nil {
		ValidatePasswordPlaintext(v, *u.Passwd.plaintext)
	}

	if u.Passwd.hash != nil {
		panic("missing password hash for user")
	}
}
