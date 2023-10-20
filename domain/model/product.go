package model

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Entity
type Product struct {
	ID          int32  `valid:"notnull"`
	Name        string `valid:"notnull"`
	Description string `valid:"notnull"`
	Price       Money  `valid:"-"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) ParseJson(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	if err = p.isValid(); err != nil {
		return err
	}

	return nil
}

func (p *Product) ToJson() ([]byte, error) {
	if err := p.isValid(); err != nil {
		return nil, err
	}

	result, err := json.Marshal(p)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func NewProduct(id int32, name string, description string, price int) (*Product, error) {
	p := &Product{
		id,
		name,
		description,
		Money{price},
	}

	if err := p.isValid(); err != nil {
		return nil, err
	}

	return p, nil
}

func NewProductFromJson(data []byte) (*Product, error) {
	p := &Product{}

	err := p.ParseJson(data)
	if err != nil {
		return nil, err
	}

	return &Product{}, nil
}
