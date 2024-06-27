package models

import "time"

type Currency int

const (
	_ = iota
	USD
	COP
)

func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case COP:
		return "COP"
	default:
		return "Unknown"
	}
}

type Account struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Entity    Entity    `json:"entity"`
	Currency  Currency  `json:"currency"`
}

type CreateAccount struct {
	Name     string   `json:"name"`
	UserID   int      `json:"user_id"`
	Entity   Entity   `json:"entity"`
	Currency Currency `json:"currency"`
}

type UpdateAccount struct {
	Name     string   `json:"name"`
	Entity   Entity   `json:"entity"`
	Currency Currency `json:"currency"`
}

func NewAccountFromParams(a *CreateAccount) *Account {
	now := time.Now()

	return &Account{
		CreatedAt: now,
		UpdatedAt: now,
		Name:      a.Name,
		UserID:    a.UserID,
		Entity:    a.Entity,
		Currency:  a.Currency,
	}
}
