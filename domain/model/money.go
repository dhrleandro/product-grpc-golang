package model

// Value Object
type Money struct {
	Value int `json:"value" valid:"int,optional"`
}

func (m *Money) ToBrazilianReal() float64 {
	return float64(m.Value) / 100
}
