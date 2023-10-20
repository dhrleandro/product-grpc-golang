package repository

import "github.com/dhrleandro/product-grpc-golang/domain/model"

type ProductRepository interface {
	Save(product *model.Product) (*model.Product, error)
	FindByName(term string) ([]*model.Product, error)
}
