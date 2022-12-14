package controllers

import (
	"fmt"
	"net/http"
	"os"
	"rsoi-lab1/models"
	"rsoi-lab1/utils"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
)

func initControllers(r *mux.Router, m *models.Models) {
	r.Use(utils.LogHandler)
	api1_r := r.PathPrefix("/api/v1/").Subrouter()

	InitPerson(api1_r, m.Person)
}

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	models := models.InitModels(db)

	initControllers(router, models)
	return router
}

func RunRouter(r *mux.Router, port uint16) error {
	c := cors.New(cors.Options{})
	handler := c.Handler(r)
	envport := os.Getenv("PORT")
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", envport), handler)
}
