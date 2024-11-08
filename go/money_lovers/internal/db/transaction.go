package db

import (
	"time"

	"money_lovers/internal/validator"
)

type Relevance int

const (
	_ = iota
	Essential
	Important
	Optional
)

func (r Relevance) String() string {
	switch r {
	case Essential:
		return "Essential"
	case Important:
		return "Important"
	case Optional:
		return "Optional"
	default:
		return "Unknown"
	}
}

type Transaction struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Date        time.Time `json:"date"`
	Concept     string    `json:"concept"`
	Description string    `json:"description"`
	Relevance   Relevance `json:"relevance"`
	AccountID   int       `json:"account_id"`
	ID          int       `json:"id"`
	Value       int32     `json:"value"`
}

type CreateTransaction struct {
	Date        time.Time `json:"date"`
	Concept     string    `json:"concept"`
	Description string    `json:"description"`
	Relevance   Relevance `json:"relevance"`
	AccountID   int       `json:"account_id"`
	Value       int32     `json:"value"`
}

type UpdateTransaction struct {
	UpdatedAt   time.Time `json:"updated_at"`
	Date        time.Time `json:"date"`
	Concept     string    `json:"concept"`
	Description string    `json:"description"`
	Relevance   Relevance `json:"relevance"`
	AccountID   int       `json:"account_id"`
	Value       int32     `json:"value"`
}

type GetTransactions struct {
	Concept     string    `json:"concept"`
	Description string    `json:"description"`
	Relevance   Relevance `json:"relevance"`
	Value       int32     `json:"value"`
	Filters
}

func NewTransactionFromParams(p *CreateTransaction) (*Transaction, error) {
	now := time.Now()

	return &Transaction{
		CreatedAt:   now,
		UpdatedAt:   now,
		Concept:     p.Concept,
		Description: p.Description,
		Value:       p.Value,
		Date:        p.Date,
		Relevance:   p.Relevance,
		AccountID:   p.AccountID,
	}, nil
}

func ValidateTransaction(v *validator.Validator, t *Transaction) {
	// Date
	v.Check(t.Date.Before(time.Now()), "date", "transaction date must not be in the future")

	// Concept
	v.Check(t.Concept != "", "concept", "must be provided")
	v.Check(len(t.Concept) >= 5, "concept", "must be at least 5 bytes long")

	// Description
	v.Check(t.Description != "", "description", "must be provided")
	v.Check(len(t.Description) >= 5, "description", "must be at least 5 bytes long")

	// Relevance
	v.Check(t.Relevance != 0, "entity", "must be provided")
	v.Check(t.Relevance > 0, "entity", "must be greater than zero")

	// AccountID
	v.Check(t.AccountID != 0, "account_id", "must be provided")

	// Value
	v.Check(t.Value >= 0, "value", "must be provided")
}
