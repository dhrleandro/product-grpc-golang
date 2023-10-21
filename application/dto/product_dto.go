package dto

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductDTO struct {
	ID          int32   `json:"id" valid:"int,required"`
	Name        string  `json:"name" valid:"notnull"`
	Description string  `json:"description" valid:"notnull"`
	Price       float64 `json:"price" valid:"float"`
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

func NewProductDTO(id int32, name string, description string, price float64) (*ProductDTO, error) {
	pdto := &ProductDTO{
		id,
		name,
		description,
		price,
	}

	var iderr error
	if pdto.ID < 0 {
		iderr = errors.New("ID: must be greater than or equal to 0")
	}

	if err := pdto.isValid(); err != nil {
		newerr := fmt.Errorf("%v;%v", err, iderr)
		return nil, newerr
	}

	if iderr != nil {
		return nil, iderr
	}

	return pdto, nil
}

func NewProductDTOFromJson(data []byte) (*ProductDTO, error) {
	pdto := &ProductDTO{}

	err := pdto.ParseJson(data)
	if err != nil {
		return nil, err
	}

	return pdto, nil
}
