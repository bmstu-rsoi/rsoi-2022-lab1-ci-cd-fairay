package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(msg)
}

func RecordNotFound(w http.ResponseWriter, recType string) {
	msg := fmt.Sprintf("Not found %s for ID", recType)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(msg)
}

func TextSuccess(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func JsonSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")

	json.NewEncoder(w).Encode(data)
}

func successCreation(w http.ResponseWriter, location string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func SuccessPersonCreation(w http.ResponseWriter, person_id int) {
	location := fmt.Sprintf("/api/v1/persons/%d", person_id)
	successCreation(w, location)
}

func SuccessPersonDeletion(w http.ResponseWriter, person_id int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(fmt.Sprintf("Person for ID %d was removed", person_id))
}

func SuccessPersonUpdate(w http.ResponseWriter, person_id int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(fmt.Sprintf("Person for ID %d was removed", person_id))
}
