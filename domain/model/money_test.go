package model_test

import (
	"testing"

	"github.com/dhrleandro/product-grpg-golang/domain/model"
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
