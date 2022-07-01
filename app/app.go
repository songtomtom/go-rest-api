package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	routes *mux.Router
	db     *gorm.DB
}

func (a *App) Init() {
	a.db = Migration()
	a.routes = Configuration(a.db)

}

func (a *App) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.routes)
}
