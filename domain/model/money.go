package model

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Value Object
type Money struct {
	Value int `valid:"int"`
}

func (m *Money) ToBrazilianReal() float32 {
	return float32(m.Value) / 100
}

func (m *Money) SetValueFromBrazilianReal(brl float32) {
	m.Value = int(brl * 100)
}
