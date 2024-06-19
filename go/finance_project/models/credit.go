package models

import "time"

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

type UpdateCreditParams struct {
	UpdatedAt    time.Time
	ClosingDate  time.Time
	DueDate      time.Time
	Identifier   string
	Entity       Entity
	Type         Type
	Rate         float32
	Total        int32
	Installments int8
}

type CreateCreditParams struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ClosingDate  time.Time
	DueDate      time.Time
	Identifier   string
	Entity       Entity
	Type         Type
	Rate         float32
	Total        int32
	Installments int8
}

type Credit struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ClosingDate  time.Time
	DueDate      time.Time
	Identifier   string
	Entity       Entity
	Type         Type
	Rate         float32
	Total        int32
	ID           int
	Installments int8
}

func NewCreditFromParams(p *CreateCreditParams) *Credit {
	now := time.Now()

	return &Credit{
		CreatedAt:    now,
		UpdatedAt:    now,
		ClosingDate:  p.ClosingDate,
		DueDate:      p.DueDate,
		Identifier:   p.Identifier,
		Entity:       p.Entity,
		Type:         p.Type,
		Rate:         p.Rate,
		Total:        p.Total,
		Installments: p.Installments,
	}
}
