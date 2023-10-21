package repository

import (
	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/mapper"
	"gorm.io/gorm"
)

// Implements github.com/dhrleandro/product-grpc-golang/domain/repository
type ProductRepositoryGORM struct {
	Db *gorm.DB
}

func (pr *ProductRepositoryGORM) Save(product *model.Product) (*model.Product, error) {
	productPersistence := mapper.ToProductPersistence(product)

	err := pr.Db.Create(productPersistence).Error
	if err != nil {
		return nil, err
	}

	p, _ := model.NewProduct(
		int32(productPersistence.ID),
		productPersistence.Name,
		productPersistence.Description,
		productPersistence.Price,
	)

	return p, nil
}
