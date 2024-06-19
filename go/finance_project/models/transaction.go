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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Concept     string
	Description string
	Value       int32
	Date        int64
	Relevance   Relevance
	Currency    Currency
	Account     Account
	ID          int
}

type CreateTransaction struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Concept     string
	Description string
	Value       int32
	Date        int64
	Relevance   Relevance
	Currency    Currency
	Account     Account
}

type UpdateTransaction struct {
	Concept     string
	Description string
	Value       int32
	Date        int64
	Relevance   Relevance
	Currency    Currency
	Account     Account
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
