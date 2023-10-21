package model

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Value Object
type Money struct {
	Value int `valid:"int"`
}

func (m *Money) ToBrazilianReal() float64 {
	return float64(m.Value) / 100
}

func (m *Money) SetValueFromBrazilianReal(brl float64) {
	m.Value = int(brl * 100)
}
