package objects

import (
	"encoding/json"
	"rsoi/lab1/errors"
)

type Person struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"not null"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

type PersonDTO struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age,omitempty" validate:"numeric"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

func (Person) TableName() string {
	return "people"
}

func (obj *Person) ToDTO() *PersonDTO {
	dto := new(PersonDTO)
	jsonStr, _ := json.Marshal(obj)
	json.Unmarshal(jsonStr, dto)
	return dto
}

func (obj *Person) Validate() error {
	if obj.Name == "" {
		return errors.InvalidPerson
	} else if obj.Age < 0 {
		return errors.InvalidPerson
	} else {
		return nil
	}
}

func (Person) ArrToDTO(src []Person) []PersonDTO {
	dst := make([]PersonDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

func (obj *PersonDTO) ToModel() *Person {
	mod := new(Person)

	jsonStr, _ := json.Marshal(obj)
	json.Unmarshal(jsonStr, mod)
	return mod
}

func (obj *PersonDTO) Validate() error {
	if obj.Name == "" {
		return errors.InvalidRequest
	} else if obj.Age < 0 {
		return errors.InvalidRequest
	} else {
		return nil
	}
}
