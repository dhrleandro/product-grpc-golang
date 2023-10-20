package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Entity
type Product struct {
	ID          int32  `json:"id" valid:"int,optional"`
	Name        string `json:"name" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	Price       Money  `json:"price" valid:"-"`
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

func NewProductFromJson(data []byte) (*Product, error) {
	p := &Product{}

	err := p.ParseJson(data)
	if err != nil {
		return nil, err
	}

	return p, nil
}
