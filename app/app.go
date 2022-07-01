package app

import (
	"fmt"
	"net/http"

	"songtomtom/rest_api/app/model"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) Init() {
	app.DB = model.Migration()
}

func (app *App) Run(port int) {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello"))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
