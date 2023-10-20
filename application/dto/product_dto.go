package dto

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type ProductDTO struct {
	ID          int32  `json:"id" valid:"int"`
	Name        string `json:"name" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	Price       int    `json:"price" valid:"int"`
}

func (pdto *ProductDTO) isValid() error {
	_, err := govalidator.ValidateStruct(pdto)
	if err != nil {
		return err
	}
	return nil
}

func (pdto *ProductDTO) ParseJson(data []byte) error {
	err := json.Unmarshal(data, pdto)
	if err != nil {
		return err
	}

	if err = pdto.isValid(); err != nil {
		return err
	}

	return nil
}

func (pdto *ProductDTO) ToJson() ([]byte, error) {
	if err := pdto.isValid(); err != nil {
		return nil, err
	}

	result, err := json.Marshal(pdto)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func NewProductDTO(id int32, name string, description string, price int) (*ProductDTO, error) {
	p := &ProductDTO{
		id,
		name,
		description,
		price,
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

func NewProductDTOFromJson(data []byte) (*ProductDTO, error) {
	pdto := &ProductDTO{}

	err := pdto.ParseJson(data)
	if err != nil {
		return nil, err
	}

	return pdto, nil
}
