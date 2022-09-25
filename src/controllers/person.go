package controllers

import (
	"rsoi/lab1/controllers/responses"
	"rsoi/lab1/errors"
	"rsoi/lab1/models"
	"rsoi/lab1/objects"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type personCtrl struct {
	model *models.PersonM
}

func InitPerson(r *mux.Router, model *models.PersonM) {
	ctrl := &personCtrl{model}
	r.HandleFunc("/persons", ctrl.add).Methods("POST")
	r.HandleFunc("/persons", ctrl.getAll).Methods("GET")

	r.HandleFunc("/persons/{id}", ctrl.get).Methods("GET")
	r.HandleFunc("/persons/{id}", ctrl.delete).Methods("DELETE")
	r.HandleFunc("/persons/{id}", ctrl.patch).Methods("PATCH")
}

func (ctrl *personCtrl) add(w http.ResponseWriter, r *http.Request) {
	person_dto := new(objects.PersonDTO)
	err := json.NewDecoder(r.Body).Decode(person_dto)
	if err != nil || person_dto.Validate() != nil {
		responses.BadRequest(w, "Invalid data")
		return
	} 

	new_person := person_dto.ToModel()
	err = ctrl.model.Create(new_person)
	switch err {
	case nil:
		responses.SuccessPersonCreation(w, new_person.Id)
	case errors.InvalidPerson:
		responses.ValidationErrorResponse(w)
	default:
		responses.BadRequest(w, "Error adding to the DB")
	}
}

func (ctrl *personCtrl) getAll(w http.ResponseWriter, r *http.Request) {
	data := ctrl.model.GetAll()
	responses.JsonSuccess(w, data)
}

func (ctrl *personCtrl) get(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Invalid id")
		return
	}

	data, err := ctrl.model.Find(id)
	switch err {
	case nil:
		responses.JsonSuccess(w, data)
	default:
		responses.RecordNotFound(w, "Person")
	}
}

func (ctrl *personCtrl) delete(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Invalid id")
		return
	}

	err = ctrl.model.Delete(id)
	switch err {
	case nil:
		responses.SuccessPersonDeletion(w, id)
	default:
		responses.RecordNotFound(w, "Person")
	}
}

func (ctrl *personCtrl) patch(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Invalid id")
		return
	}

	person_dto := new(objects.PersonDTO)
	err = json.NewDecoder(r.Body).Decode(person_dto)
	if err != nil {
		responses.BadRequest(w, "Invalid data")
		return
	}

	person := person_dto.ToModel()
	person.Id = id

	upd_person, err := ctrl.model.Patch(person)
	switch err {
	case nil:
		responses.JsonSuccess(w, upd_person)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, "Person")
	default:
		responses.BadRequest(w, "Invalid data")
	}
}
