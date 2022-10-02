package models

import (
	"rsoi-lab1/errors"
	"rsoi-lab1/objects"
	"rsoi-lab1/repository"
)

type PersonM struct {
	rep    repository.PersonRep
	models *Models
}

func NewPersonM(rep repository.PersonRep, models *Models) *PersonM {
	return &PersonM{rep, models}
}

func (model *PersonM) Create(obj *objects.Person) error {
	if err := obj.Validate(); err != nil {
		return err
	}
	if err := model.rep.Create(obj); err != nil {
		return errors.DBAdditionError
	}

	return nil
}

func (model *PersonM) GetAll() []objects.Person {
	return model.rep.GetAll()
}

func (model *PersonM) Find(id int) (*objects.Person, error) {
	person, err := model.rep.Find(id)
	if err != nil {
		return nil, errors.RecordNotFound
	} else {
		return person, nil
	}
}

func (model *PersonM) Delete(id int) error {
	return model.rep.Delete(id)
}

func (model *PersonM) Patch(obj *objects.Person) (*objects.Person, error) {
	return model.rep.Patch(obj)
}
