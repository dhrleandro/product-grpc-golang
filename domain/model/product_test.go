package model_test

import (
	"testing"

	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	var id int32 = 10
	name := "Table"
	description := "wood"
	price := 2535
	p, err := model.NewProduct(id, name, description, price)
	require.Equal(t, err, nil)
	require.Equal(t, p.ID, id)
	require.Equal(t, p.Name, name)
	require.Equal(t, p.Description, description)
	require.Equal(t, p.Price.Value, price)

	id = 0
	name = ""
	description = ""
	price = 0
	_, err = model.NewProduct(id, name, description, price)
	require.NotEqual(t, err, nil)

	id = -58
	name = "Table"
	description = "wood"
	price = 2535
	_, err = model.NewProduct(id, name, description, price)
	require.NotEqual(t, err, nil)
}
