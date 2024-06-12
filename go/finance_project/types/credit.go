package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreditBase struct {
	ClosingDate         time.Time          `bson:"closingDate,omitempty" json:"closingDate,omitempty"`
	DueDate             time.Time          `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
	Entity              string             `bson:"entity,omitempty" json:"entity,omitempty"`
	Identifier          string             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                string             `bson:"type,omitempty" json:"type,omitempty"`
	Rate                float32            `bson:"rate,omitempty" json:"rate,omitempty"`
	Total               int16              `bson:"total,omitempty" json:"total,omitempty"`
	ID                  primitive.ObjectID `bson:"_id" json:"id"`
	Number_Installments int8               `bson:"installments,omitempty" json:"installments,omitempty"`
}

type UpdateCreditParams struct {
	CreditBase
}

type CreateCreditParams struct {
	CreditBase
}

type Credit struct {
	CreditBase

	ID primitive.ObjectID
}

func NewCreditFromParams(p *CreateCreditParams) (*Credit, error) {
	return &Credit{
		CreditBase: p.CreditBase,
	}, nil
}
