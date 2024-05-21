package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

type UpdateTransactionParams struct {
	Concept     string `json:"concept"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
	Date        int64  `json:"date"`
	Reference   string `json:"reference"`
	Category    string `bson:"category" json:"category"`
	Account     string `bson:"account" json:"account"`
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
	} else if len(up.Reference) > 0 {
		m["Reference"] = up.Reference
	} else if len(up.Category) > 0 {
		m["Category"] = up.Category
	} else if len(up.Account) > 0 {
		m["Account"] = up.Account
	}

	return m
}

type CreateTransactionParams struct {
	Concept     string `json:"concept"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
	Date        int64  `json:"date"`
	Status      string `json:"status"`
	Currency    string `json:"currency"`
	Account     string `bson:"account" json:"account"`
}

type TransactionParams struct {
	Concept     string `json:"concept"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
	Date        int64  `json:"date"`
	Status      string `json:"status"`
	Currency    string `json:"currency"`
	Account     string `bson:"account" json:"account"`
}

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Concept     string             `bson:"concept,omitempty" json:"concept,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Value       int32              `bson:"value,omitempty" json:"value,omitempty"`
	Date        int64              `bson:"date,omitempty" json:"date,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
	Currency    string             `bson:"currency,omitempty" json:"currency,omitempty"`
	Account     string             `bson:"account,omitempty" json:"account,omitempty"`
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

	if len(tp.Account) < minAccountLen {
		return fmt.Errorf("account length should me larger than %d", minAccountLen)
	}

	return nil
}

func NewTransactionFromParams(p CreateTransactionParams) (*Transaction, error) {
	return &Transaction{
		Concept:     p.Concept,
		Description: p.Description,
		Value:       p.Value,
		Date:        p.Date,
		Status:      p.Status,
		Currency:    p.Currency,
		Account:     p.Account,
	}, nil
}
