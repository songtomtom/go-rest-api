package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Configuration(db *gorm.DB) *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	// routes.HandleFunc("/user", handler.FindAll(db))
	return routes
}
