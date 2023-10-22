package dto_test

import (
	"testing"

	"github.com/dhrleandro/product-grpc-golang/application/dto"
	"github.com/stretchr/testify/require"
)

func TestNewProductDTO(t *testing.T) {
	var id int32 = 10
	name := "Table"
	description := "wood"
	var price float32 = 253.5
	pdto, err := dto.NewProductDTO(id, name, description, price)

	require.Equal(t, err, nil)
	require.Equal(t, pdto.ID, id)
	require.Equal(t, pdto.Name, name)
	require.Equal(t, pdto.Description, description)
	require.Equal(t, pdto.Price, price)

	id = 0
	name = ""
	description = ""
	price = 0
	_, err = dto.NewProductDTO(id, name, description, price)
	require.NotEqual(t, err, nil)

	id = -159
	name = ""
	description = ""
	price = 0
	_, err = dto.NewProductDTO(id, name, description, price)
	require.NotEqual(t, err, nil)
}

func TestProductToJsonDTO(t *testing.T) {
	p, _ := dto.NewProductDTO(5, "Computer", "Core i7", 534.8)
	json, err := p.ToJson()

	require.Equal(t, err, nil)

	testJson := []byte(`{"id":5,"name":"Computer","description":"Core i7","price":534.8}`)
	require.Equal(t, json, testJson)
}

func TestNewProductFromJsonDTO(t *testing.T) {
	json := []byte(`{"id":58,"name":"Computer","description":"Core i7","price":100.38}`)
	p, err := dto.NewProductDTOFromJson(json)

	require.Equal(t, err, nil)
	require.Equal(t, p.ID, int32(58))
	require.Equal(t, p.Name, "Computer")
	require.Equal(t, p.Description, "Core i7")
	require.Equal(t, p.Price, float32(100.38))
}
