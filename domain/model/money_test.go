package model_test

import (
	"testing"

	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/stretchr/testify/require"
)

func TestToBrazilianReal(t *testing.T) {
	m := &model.Money{25698}
	require.Equal(t, m.ToBrazilianReal(), float32(256.98))

	m = &model.Money{1}
	require.Equal(t, m.ToBrazilianReal(), float32(0.01))

	m = &model.Money{32049}
	require.Equal(t, m.ToBrazilianReal(), float32(320.49))
}

func TestSetValueFromBrazilianReal(t *testing.T) {
	m := &model.Money{}
	var brl float32 = 236.89

	m.SetValueFromBrazilianReal(brl)
	require.Equal(t, m.Value, 23689)
}
