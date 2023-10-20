package mapper

import (
	"github.com/dhrleandro/product-grpc-golang/application/dto"
	"github.com/dhrleandro/product-grpc-golang/domain/model"
)

func ToProductDomain(pdto *dto.ProductDTO) (*model.Product, error) {
	return nil, nil
}

func ToProductDTO(product *model.Product) (*dto.ProductDTO, error) {
	return nil, nil
}

func ToProductPersistence(product *model.Product) (*model.Product, error) {
	return nil, nil
}
