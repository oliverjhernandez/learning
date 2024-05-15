package types

type Transaction struct {
	ID      string `bson:"_id,omitempty" json:"id,omitempty"`
	Concept string `bsin:"concept" json:"concept"`
	Value   int32  `bsin:"value" json:"value"`
}
