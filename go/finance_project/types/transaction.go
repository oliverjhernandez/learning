package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	bcryptCost        = 12
	minConceptLen     = 5
	minDescriptionLen = 1
	minValue          = 1
	minReferenteLen   = 1
	minAccountLen     = 3
)

type TransactionParams struct {
	Concept     string `json:"concept"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
	Date        int64  `json:"date"`
	Reference   string `json:"reference"`
	Category    string `bson:"category" json:"category"`
	Account     string `bson:"account" json:"account"`
}

type Transaction struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TransactionParams
}

func (tp TransactionParams) Validate() error {
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

	if len(tp.Account) < minAccountLen {
		return fmt.Errorf("account length should me larger than %d", minAccountLen)
	}

	return nil
}

func NewTransactionFromParams(p TransactionParams) (*Transaction, error) {
	return &Transaction{
		TransactionParams: TransactionParams{
			Concept:     p.Concept,
			Description: p.Description,
			Value:       p.Value,
			Date:        p.Date,
			Reference:   p.Reference,
			Category:    p.Category,
			Account:     p.Account,
		},
	}, nil
}
