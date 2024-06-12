package types

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type UserBase struct {
	FirstName string `bson:"firstname,omitempty" json:"firstname,omitempty"`
	LastName  string `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email     string `bson:"email,omitempty" json:"email,omitempty"`
}

type CreateUserParams struct {
	UserBase
	Passwd string `bson:"passwd,omitempty"`
}

type UpdateUserParams struct {
	UserBase
}

type User struct {
	UserBase
	Passwd string             `bson:"passwd,omitempty"`
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
}

func (up UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(up.FirstName) > 0 {
		m["firstName"] = up.FirstName
	} else if len(up.LastName) > 0 {
		m["lastName"] = up.LastName
	} else if len(up.Email) > 0 {
		m["email"] = up.Email
	}

	return m
}

func NewUserFromParams(p CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(p.Passwd), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		UserBase: p.UserBase,
		Passwd:   string(encpw),
	}, nil
}
