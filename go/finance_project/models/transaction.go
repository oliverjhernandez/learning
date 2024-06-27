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

type Account int

const (
	_ = iota
	SAVINGS
	CHECKINGS
	LOAN
)

func (a Account) String() string {
	switch a {
	case SAVINGS:
		return "SAVINGS"
	case CHECKINGS:
		return "CHECKINGS"
	case LOAN:
		return "LOAN"
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

func NewTransactionFromParams(p CreateTransaction) *Transaction {
	now := time.Now()

	return &Transaction{
		CreatedAt:   now,
		UpdatedAt:   now,
		Concept:     p.Concept,
		Description: p.Description,
		Value:       p.Value,
		Date:        p.Date,
		Relevance:   p.Relevance,
		Currency:    p.Currency,
		Account:     p.Account,
	}
}
