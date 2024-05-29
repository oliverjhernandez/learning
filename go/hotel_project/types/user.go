package types

import (
	"fmt"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswdLen    = 12
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (up UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(up.FirstName) > 0 {
		m["firstName"] = up.FirstName
	}
	if len(up.LastName) > 0 {
		m["lastName"] = up.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Passwd    string `json:"passwd"`
}

func (p CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(p.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(p.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
	}
	if len(p.Passwd) < minPasswdLen {
		errors["passwd"] = fmt.Sprintf("passwd length should be at least %d characters", minPasswdLen)
	}
	if !isValidEmail(p.Email) {
		errors["email"] = fmt.Sprintf("email %s is invalid", p.Email)
	}
	return errors
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       string             `bson:"firstName" json:"firstName"`
	LastName        string             `bson:"lastName" json:"lastName"`
	Email           string             `bson:"email" json:"email"`
	EncryptedPasswd string             `bson:"encryptedPasswd" json:"-"`
}

func IsValidPasswd(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Passwd), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:       params.FirstName,
		LastName:        params.LastName,
		Email:           params.Email,
		EncryptedPasswd: string(encpw),
	}, nil
}
