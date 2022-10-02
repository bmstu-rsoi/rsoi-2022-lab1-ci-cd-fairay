package unit_test

import (
	"rsoi-lab1/errors"
	"rsoi-lab1/mocks"
	"rsoi-lab1/models"
	"rsoi-lab1/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	t.Run("correct operation", func(t *testing.T) {
		mockRep := new(mocks.PersonRep)
		model := models.NewPersonM(mockRep, nil)
		obj := &objects.Person{
			Id: 0, Name: "Celestino.Murazik88", Age: 31,
			Address: "9124 Jacobi Flats", Work: "Beatty - Cronin",
		}

		mockRep.On("Create", obj).Return(nil).Once()

		err := model.Create(obj)

		assert.Nil(t, err, "Create person have unexpected error")
		mockRep.AssertExpectations(t)
	})

	t.Run("invalid name", func(t *testing.T) {
		mockRep := new(mocks.PersonRep)

		model := models.NewPersonM(mockRep, nil)
		obj := &objects.Person{
			Id: 0, Age: 31,
			Address: "9124 Jacobi Flats", Work: "Beatty - Cronin",
		}

		err := model.Create(obj)

		assert.Equal(t, errors.InvalidPerson, err, "Create person have unexpected error")
		mockRep.AssertExpectations(t)
	})
}

func TestFind(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		mockRep := new(mocks.PersonRep)
		model := models.NewPersonM(mockRep, nil)
		obj := &objects.Person{
			Id: 10, Name: "Celestino.Murazik88", Age: 31,
			Address: "9124 Jacobi Flats", Work: "Beatty - Cronin",
		}

		mockRep.On("Find", obj.Id).Return(obj, nil).Once()

		res, err := model.Find(obj.Id)

		assert.Nil(t, err, "Find person have unexpected error")
		assert.Equal(t, obj, res, "Unexpected object found")
		mockRep.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockRep := new(mocks.PersonRep)
		model := models.NewPersonM(mockRep, nil)

		mockRep.On("Find", 0).Return(nil, errors.RecordNotFound).Once()

		res, err := model.Find(0)

		var nil_person *objects.Person = nil
		assert.Equal(t, errors.RecordNotFound, err, "Find person have unexpected error")
		assert.Equal(t, nil_person, res, "Unexpected object found")
		mockRep.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deleted", func(t *testing.T) {
		mockRep := new(mocks.PersonRep)
		model := models.NewPersonM(mockRep, nil)
		obj := &objects.Person{
			Id: 10, Name: "Celestino.Murazik88", Age: 31,
			Address: "9124 Jacobi Flats", Work: "Beatty - Cronin",
		}

		mockRep.On("Delete", obj.Id).Return(nil).Once()

		err := model.Delete(obj.Id)

		assert.Nil(t, err, "Delete person have unexpected error")
		mockRep.AssertExpectations(t)
	})
}

func TestPatch(t *testing.T) {
	t.Run("patched", func(t *testing.T) {
		mockRep := new(mocks.PersonRep)
		model := models.NewPersonM(mockRep, nil)
		patch := &objects.Person{
			Age:     31,
			Address: "9124 Jacobi Flats",
		}
		upd_obj := &objects.Person{
			Id: 10, Name: "Celestino.Murazik88", Age: 31,
			Address: "9124 Jacobi Flats", Work: "Beatty - Cronin",
		}

		mockRep.On("Patch", patch).Return(upd_obj, nil).Once()

		res, err := model.Patch(patch)

		assert.Nil(t, err, "Patch person have unexpected error")
		assert.Equal(t, upd_obj, res, "Unexpected object after patch")
		mockRep.AssertExpectations(t)
	})
}
