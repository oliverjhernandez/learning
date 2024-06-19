package models

import "time"

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Email     string
	Passwd    string
	ID        int
}

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
	Passwd    string
}

type UpdateUser struct {
	FirstName string
	LastName  string
	Email     string
}

func NewUserFromParams(u CreateUser) *User {
	now := time.Now()

	return &User{
		CreatedAt: now,
		UpdatedAt: now,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
