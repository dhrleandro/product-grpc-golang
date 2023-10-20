package model

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Entity
//
// For this project this entity is anemic model but let's assume it is not
type Product struct {
	ID          int32  `valid:"int,optional"`
	Name        string `valid:"notnull"`
	Description string `valid:"notnull"`
	Price       Money  `valid:"required"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(id int32, name string, description string, price int) (*Product, error) {
	p := &Product{
		id,
		name,
		description,
		Money{price},
	}

	var iderr error
	if p.ID < 0 {
		iderr = errors.New("ID: must be greater than or equal to 0")
	}

	if err := p.isValid(); err != nil {
		newerr := fmt.Errorf("%v;%v", err, iderr)
		return nil, newerr
	}

	if iderr != nil {
		return nil, iderr
	}

	return p, nil
}
