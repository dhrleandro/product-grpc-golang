package repository

import (
	"github.com/dhrleandro/product-grpc-golang/domain/model"
	dbmodel "github.com/dhrleandro/product-grpc-golang/infrastructure/database/model"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/mapper"
	"gorm.io/gorm"
)

// Implements github.com/dhrleandro/product-grpc-golang/domain/repository
type ProductRepositoryGORM struct {
	Db *gorm.DB
}

func (pr *ProductRepositoryGORM) Save(product *model.Product) (*model.Product, error) {
	productPersistence := mapper.ToProductPersistence(product)

	err := pr.Db.Save(productPersistence).Error
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

func (pr *ProductRepositoryGORM) FindByName(term string) ([]*model.Product, error) {
	var products []dbmodel.Product
	pr.Db.Where("name LIKE ?", "%"+term+"%").Find(&products)

	results := []*model.Product{}
	for _, v := range products {

		p, _ := model.NewProduct(
			int32(v.ID),
			v.Name,
			v.Description,
			v.Price,
		)
		results = append(results, p)
	}

	return results, nil
}
