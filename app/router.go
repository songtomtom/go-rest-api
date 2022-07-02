package app

import (
	"songtomtom/rest_api/app/handler"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Configuration(db *gorm.DB) *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", handler.Index())
	routes.HandleFunc("/user", handler.FindAll(db))
	return routes
}
