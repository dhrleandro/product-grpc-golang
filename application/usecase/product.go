package usecase

import (
	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/dhrleandro/product-grpc-golang/domain/repository"
)

type ProductUseCase struct {
	Pr repository.ProductRepository
}

func (cpuc *ProductUseCase) CreateProduct(
	name string,
	description string,
	price float64,
) (*model.Product, error) {
	m := &model.Money{}
	m.SetValueFromBrazilianReal(price)
	product, err := model.NewProduct(0, name, description, m.Value)
	if err != nil {
		return nil, err
	}

	ret, err := cpuc.Pr.Save(product)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (cpuc *ProductUseCase) FindProduct(term string) ([]*model.Product, error) {

	ret, err := cpuc.Pr.FindByName(term)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
