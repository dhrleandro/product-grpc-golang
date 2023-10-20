package model_test

import (
	"testing"

	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/stretchr/testify/require"
)

func TestToBrazilianReal(t *testing.T) {
	m := &model.Money{25698}
	require.Equal(t, m.ToBrazilianReal(), float64(256.98))
	require.Equal(t, m.ToBrazilianReal(), 256.98)

	m = &model.Money{1}
	require.Equal(t, m.ToBrazilianReal(), float64(0.01))
	require.Equal(t, m.ToBrazilianReal(), 0.01)
}

func TestSetValueFromBrazilianReal(t *testing.T) {
	m := &model.Money{}
	var brl float64 = 236.89

	m.SetValueFromBrazilianReal(brl)
	require.Equal(t, m.Value, 23689)
}
