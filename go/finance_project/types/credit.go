package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity int

const (
	_ = iota
	BANCOLOMBIA
	AV_VILLAS
	DAVIVIENDA
)

func (e Entity) String() string {
	switch e {
	case BANCOLOMBIA:
		return "BANCOLOMBIA"
	case AV_VILLAS:
		return "AV VILLAS"
	case DAVIVIENDA:
		return "DAVIVIENDA"
	default:
		return "Unknown"
	}
}

type Type int

const (
	_ = iota
	LIBRE_INVERSION
	HIPOTECARIO
	VEHICULO
)

func (t Type) String() string {
	switch t {
	case LIBRE_INVERSION:
		return "LIBRE INVERSION"
	case HIPOTECARIO:
		return "HIPOTECARIO"
	case VEHICULO:
		return "VEHICULO"
	default:
		return "UNKNOWN"
	}
}

type CreditBase struct {
	ClosingDate         time.Time          `bson:"closingDate,omitempty" json:"closingDate,omitempty"`
	DueDate             time.Time          `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
	Identifier          string             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Entity              Entity             `bson:"entity,omitempty" json:"entity,omitempty"`
	Type                Type               `bson:"type,omitempty" json:"type,omitempty"`
	Rate                float32            `bson:"rate,omitempty" json:"rate,omitempty"`
	Total               int32              `bson:"total,omitempty" json:"total,omitempty"`
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
