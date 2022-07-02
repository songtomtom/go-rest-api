package app

import (
	"songtomtom/rest_api/app/handler"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Configuration(db *gorm.DB) *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", handler.Index())
	routes.HandleFunc("/user", handler.FindAll(db)).Methods("GET")
	routes.HandleFunc("/user", handler.Create(db)).Methods("POST")
	routes.HandleFunc("/user/{id}", handler.FindById(db)).Methods("GET")
	routes.HandleFunc("/user/{id}", handler.Update(db)).Methods("PUT")
	routes.HandleFunc("/user/{id}", handler.Delete(db)).Methods("DELETE")
	return routes
}
