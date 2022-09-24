package repository

import (
	"rsoi/lab1/errors"
	"rsoi/lab1/objects"

	"github.com/jinzhu/gorm"
)

type PersonRep interface {
	Create(obj *objects.Person) error
	GetAll() []objects.Person
	Find(id int) (*objects.Person, error)
	Delete(id int) error
	Patch(obj *objects.Person) (*objects.Person, error)
}

type PGPersonRep struct {
	db *gorm.DB
}

func NewPersonRep(db *gorm.DB) *PGPersonRep {
	return &PGPersonRep{db}
}

func (rep *PGPersonRep) Create(obj *objects.Person) error {
	return rep.db.Create(obj).Error
}

func (rep *PGPersonRep) GetAll() []objects.Person {
	temp := []objects.Person{}
	rep.db.Find(&temp)
	return temp
}

func (rep *PGPersonRep) Find(id int) (*objects.Person, error) {
	temp := new(objects.Person)
	err := rep.db.Where("id = ?", id).First(temp).Error
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		temp, err = nil, errors.RecordNotFound
	default:
		temp, err = nil, errors.UnknownError
	}

	return temp, err
}

func (rep *PGPersonRep) Delete(id int) error {
	person := objects.Person{Id: id}
	return rep.db.Delete(person).Error
}

func (rep *PGPersonRep) Patch(obj *objects.Person) (*objects.Person, error) {
	res := rep.db.Model(obj).Updates(obj)
	if res.Error != nil {
		return nil, errors.UnknownError
	} else if res.RowsAffected != 1 {
		return nil, errors.RecordNotFound
	}

	return rep.Find(obj.Id)
}
