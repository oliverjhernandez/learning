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
	LIBRE
	HIPOTECARIO
	VEHICULO
)

func (t Type) String() string {
	switch t {
	case LIBRE:
		return "LIBRE"
	case HIPOTECARIO:
		return "HIPOTECARIO"
	case VEHICULO:
		return "VEHICULO"
	default:
		return "UNKNOWN"
	}
}

type UpdateCredit struct {
	UpdatedAt    time.Time `json:"updated_at"`
	ClosingDate  time.Time `json:"closing_date"`
	DueDate      time.Time `json:"due_date"`
	Identifier   string    `json:"identifier"`
	Entity       Entity    `json:"entity"`
	Type         Type      `json:"type"`
	Rate         float32   `json:"rate"`
	Total        int32     `json:"total"`
	Installments int8      `json:"installments"`
	UserID       int       `json:"user_id"`
}

type CreateCredit struct {
	ClosingDate  time.Time `json:"closing_date"`
	DueDate      time.Time `json:"due_date"`
	Identifier   string    `json:"identifier"`
	Entity       Entity    `json:"entity"`
	Type         Type      `json:"type"`
	Rate         float32   `json:"rate"`
	Total        int32     `json:"total"`
	Installments int8      `json:"installments"`
	UserID       int       `json:"user_id"`
}

type Credit struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosingDate  time.Time `json:"closing_date"`
	DueDate      time.Time `json:"due_date"`
	Identifier   string    `json:"identifier"`
	Entity       Entity    `json:"entity"`
	Type         Type      `json:"type"`
	Rate         float32   `json:"rate"`
	Total        int32     `json:"total"`
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Installments int8      `json:"installments"`
}

func NewCreditFromParams(p *CreateCredit) *Credit {
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
		UserID:       p.UserID,
	}
}
