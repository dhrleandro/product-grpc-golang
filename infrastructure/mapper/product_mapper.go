package mapper

import (
	"time"

	"github.com/dhrleandro/product-grpc-golang/application/dto"
	"github.com/dhrleandro/product-grpc-golang/domain/model"
	dbmodel "github.com/dhrleandro/product-grpc-golang/infrastructure/database/model"
)

func ToProductDomain(pdto *dto.ProductDTO) (*model.Product, error) {
	m := &model.Money{}
	m.SetValueFromBrazilianReal(pdto.Price)
	product, err := model.NewProduct(
		pdto.ID,
		pdto.Name,
		pdto.Description,
		m.Value,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func ToProductDTO(product *model.Product) (*dto.ProductDTO, error) {
	pdto, err := dto.NewProductDTO(
		product.ID,
		product.Name,
		product.Description,
		product.Price.ToBrazilianReal(),
	)

	if err != nil {
		return nil, err
	}

	return pdto, nil
}

func ToProductPersistence(product *model.Product) *dbmodel.Product {
	dbProduct := &dbmodel.Product{
		ID:          uint(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price.Value,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	return dbProduct
}
