package model

// Value Object
type Money struct {
	Value int `valid:"int,optional"`
}

func (m *Money) ToBrazilianReal() float64 {
	return float64(m.Value) / 100
}
