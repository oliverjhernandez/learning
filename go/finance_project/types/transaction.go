package types

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minConceptLen     = 5
	minDescriptionLen = 1
	minValue          = 1
)

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

func (r Relevance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Relevance) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "Essential":
		*r = Essential
	case "Important":
		*r = Important
	case "Optional":
		*r = Optional
	default:
		return fmt.Errorf("invalid relevance value: %s", s)
	}
	return nil
}

type UpdateTransactionParams struct {
	TransactionBase
}

func (up UpdateTransactionParams) ToBSON() bson.M {
	m := bson.M{}
	if len(up.Concept) > 0 {
		m["Concept"] = up.Concept
	} else if len(up.Description) > 0 {
		m["Description"] = up.Description
	} else if up.Value > 0 {
		m["Value"] = up.Value
	} else if up.Date > 0 {
		m["Date"] = up.Date
	}

	return m
}

type Level string

type TransactionBase struct {
	Concept     string    `json:"concept"`
	Description string    `json:"description"`
	Value       int32     `json:"value"`
	Date        int64     `json:"date"`
	Relevance   Relevance `json:"Relevance"`
	Currency    Currency  `json:"currency"`
	Account     Account   `bson:"account" json:"account"`
}

type CreateTransactionParams struct {
	TransactionBase
}

type Transaction struct {
	TransactionBase
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
}

func (tp CreateTransactionParams) Validate() error {
	if len(tp.Concept) < minConceptLen {
		return fmt.Errorf("concept length should me larger than %d", minConceptLen)
	}

	if len(tp.Description) < minDescriptionLen {
		return fmt.Errorf("description length should me larger than %d", minDescriptionLen)
	}

	if tp.Value < minValue {
		return fmt.Errorf("value should me larger than %d", minValue)
	}

	if tp.Date > time.Now().Unix() {
		return fmt.Errorf("date should me smaller than current date")
	}

	return nil
}

func NewTransactionFromParams(p CreateTransactionParams) (*Transaction, error) {
	return &Transaction{
		TransactionBase: p.TransactionBase,
	}, nil
}
