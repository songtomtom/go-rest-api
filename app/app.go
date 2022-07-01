package app

import (
	"fmt"
	"net/http"

	"songtomtom/rest_api/app/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	routers *mux.Router
	db      *gorm.DB
}

func (a *App) Init() {
	a.db = model.Migration()
	routers := mux.NewRouter()
	routers.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	a.routers = routers
}

func (a *App) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.routers)
}
