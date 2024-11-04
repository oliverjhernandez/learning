package db

import (
	"time"

	"casita/internal/validator"
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

type Credit struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosingDay   int8      `json:"closing_day"`
	DueDay       int8      `json:"due_day"`
	Identifier   string    `json:"identifier"`
	Entity       Entity    `json:"entity"`
	Type         Type      `json:"type"`
	Rate         float32   `json:"rate"`
	Total        int32     `json:"total"`
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Installments int8      `json:"installments"`
}

type UpdateCredit struct {
	UpdatedAt    time.Time `json:"updated_at"`
	ClosingDay   int8      `json:"closing_day"`
	DueDay       int8      `json:"due_day"`
	Identifier   string    `json:"identifier"`
	Entity       Entity    `json:"entity"`
	Type         Type      `json:"type"`
	Rate         float32   `json:"rate"`
	Total        int32     `json:"total"`
	Installments int8      `json:"installments"`
	UserID       int       `json:"user_id"`
}

type CreateCredit struct {
	ClosingDay   int8    `json:"closing_day"`
	DueDay       int8    `json:"due_day"`
	Identifier   string  `json:"identifier"`
	Entity       Entity  `json:"entity"`
	Type         Type    `json:"type"`
	Rate         float32 `json:"rate"`
	Total        int32   `json:"total"`
	Installments int8    `json:"installments"`
	UserID       int     `json:"user_id"`
}

func NewCreditFromParams(p *CreateCredit) (*Credit, error) {
	now := time.Now()

	return &Credit{
		CreatedAt:    now,
		UpdatedAt:    now,
		ClosingDay:   p.ClosingDay,   // Billing cycle end date
		DueDay:       p.DueDay,       // Limit date to pay at least a min
		Identifier:   p.Identifier,   // Credit identifier, usually provideed by the entity
		Entity:       p.Entity,       // Bank, Person, any institution that granted the credit
		Type:         p.Type,         // enum makes it self-explenatory
		Rate:         p.Rate,         // Interest rate
		Total:        p.Total,        // Total granted for the credit
		Installments: p.Installments, // Top number of payments to finalize the credit
		UserID:       p.UserID,       // Owner of the credit
	}, nil
}

func ValidateCredit(v *validator.Validator, c *Credit) {
	// ClosingDate
	v.Check(c.ClosingDay > 0, "closing_day", "must be greater than zero")
	v.Check(c.ClosingDay <= 30, "closing_day", "must be lower or equal to 30")
	v.Check(c.ClosingDay <= c.DueDay, "closing_day", "must take place before due_day")

	// DueDate
	v.Check(c.DueDay > 0, "due_day", "must be greater than zero")
	v.Check(c.DueDay <= 30, "due_day", "must be lower or equal to zero")
	v.Check(c.DueDay >= c.ClosingDay, "due_date", "must take place after closing_day")

	// Identifier
	v.Check(c.Identifier != "", "identifier", "must be provided")
	v.Check(len(c.Identifier) >= 5, "identifier", "must be at least 5 bytes long")

	// Entity
	v.Check(c.Entity != 0, "entity", "must be provided")

	// Type
	v.Check(c.Type != 0, "passwd", "must be provided")

	// Rate
	v.Check(c.Rate >= 0, "rate", "must be provided")

	// Total
	v.Check(c.Total > 0, "total", "must be greater than zero")

	// Installments
	v.Check(c.Installments > 0, "installments", "must be greater than zero")

	// UserID
	v.Check(c.UserID != 0, "user_id", "must be assigned to a user")
}
