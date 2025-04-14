package data

// #TODO: Create CRUD for Entity
type Entity struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Kind     Kind     `json:"kind"`
	Provider Provider `json:"provider"`
	Currency Currency `json:"currency"`
}

type Kind int

const (
	_ = iota
	WALLET
	CARD
	SAVINGS
	CHECKINGS
)

type Provider int

const (
	_ = iota
	BANCOLOMBIA
	AV_VILLAS
	DAVIVIENDA
)

func (e Provider) String() string {
	switch e {
	case BANCOLOMBIA:
		return "Bancolombia"
	case AV_VILLAS:
		return "Citi"
	case DAVIVIENDA:
		return "BDV"
	default:
		return "Unknown"
	}
}

type Currency int

const (
	_ = iota
	USD
	BSF
	COP
)

func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case BSF:
		return "BSF"
	case COP:
		return "COP"
	default:
		return "Unknown"
	}
}
