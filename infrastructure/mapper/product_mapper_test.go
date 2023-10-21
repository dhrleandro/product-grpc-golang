package mapper_test

// https://softwareengineering.stackexchange.com/questions/429957/what-is-the-responsibility-of-infrastructure-layer-in-a-clean-architecture

import (
	"testing"

	"github.com/dhrleandro/product-grpc-golang/application/dto"
	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/mapper"
	"github.com/stretchr/testify/require"
)

func TestToProductDomain(t *testing.T) {
	pdto, err := dto.NewProductDTO(
		5,
		"Table",
		"Wood",
		5847.99,
	)
	require.Equal(t, err, nil)

	p, err := mapper.ToProductDomain(pdto)
	require.Equal(t, err, nil)

	require.Equal(t, p.ID, pdto.ID)
	require.Equal(t, p.Name, pdto.Name)
	require.Equal(t, p.Description, pdto.Description)
	require.Equal(t, p.Price.ToBrazilianReal(), pdto.Price)
}

func TestToProductDTO(t *testing.T) {
	p, err := model.NewProduct(
		5,
		"Table",
		"Wood",
		584799,
	)
	require.Equal(t, err, nil)

	pdto, err := mapper.ToProductDTO(p)
	require.Equal(t, err, nil)

	require.Equal(t, pdto.ID, p.ID)
	require.Equal(t, pdto.Name, p.Name)
	require.Equal(t, pdto.Description, p.Description)
	require.Equal(t, pdto.Price, p.Price.ToBrazilianReal())
}

func TestToPersistence(t *testing.T) {
	p, err := model.NewProduct(
		5,
		"Table",
		"Wood",
		584799,
	)
	require.Equal(t, err, nil)

	dbProduct := mapper.ToProductPersistence(p)

	require.Equal(t, dbProduct.ID, uint(p.ID))
	require.Equal(t, dbProduct.Name, p.Name)
	require.Equal(t, dbProduct.Description, p.Description)
	require.Equal(t, dbProduct.Price, p.Price.Value)
}
