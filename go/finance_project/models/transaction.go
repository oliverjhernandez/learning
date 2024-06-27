package models

import "time"

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

func NewTransactionFromParams(p CreateTransaction) (*Transaction, error) {
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
