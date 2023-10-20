package model_test

import (
	"testing"

	"github.com/dhrleandro/product-grpg-golang/domain/model"
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
}

func TestProductToJson(t *testing.T) {
	p, _ := model.NewProduct(5, "Computer", "Core i7", 5348)
	json, err := p.ToJson()

	require.Equal(t, err, nil)

	testJson := []byte(`{"id":5,"name":"Computer","description":"Core i7","price":{"value":5348}}`)
	require.Equal(t, json, testJson)
}

func TestNewProductFromJson(t *testing.T) {
	json := []byte(`{"id":58,"name":"Computer","description":"Core i7","price":{"value":100}}`)
	p, err := model.NewProductFromJson(json)

	require.Equal(t, err, nil)
	require.Equal(t, p.ID, int32(58))
	require.Equal(t, p.Name, "Computer")
	require.Equal(t, p.Description, "Core i7")
	require.Equal(t, p.Price.Value, 100)
}
