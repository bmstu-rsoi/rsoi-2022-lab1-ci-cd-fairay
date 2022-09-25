package models

import (
	"rsoi/lab1/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Person *PersonM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Person = NewPersonM(repository.NewPersonRep(db), models)

	return models
}
