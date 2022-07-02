package app

import (
	"fmt"
	"net/http"
	"songtomtom/rest_api/app/handler"
	"songtomtom/rest_api/app/model"
	"songtomtom/rest_api/config"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	routes *mux.Router
	db     *gorm.DB
}

type ResponseHandlerFunc func(db *gorm.DB, w http.ResponseWriter, r *http.Request) func(w http.ResponseWriter, r *http.Request)

func (app *App) Run(port int) {
	app.Migration()
	app.Configuration()
	fmt.Println("Server Start")
	http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes)

}

func (app *App) Migration() {
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connect error.")
	}
	db.AutoMigrate(&model.User{})
	app.db = db
}

func (app *App) Configuration() {
	app.routes = mux.NewRouter()
	app.Get("/", app.ResponseHandler(handler.Index))
	app.Get("/user", app.ResponseHandler(handler.FindAll))
	app.Post("/user", app.ResponseHandler(handler.Create))
	app.Get("/user/{id}", app.ResponseHandler(handler.FindById))
	app.Put("/user/{id}", app.ResponseHandler(handler.Update))
	app.Delete("/user/{id}", app.ResponseHandler(handler.Delete))
}

func (app *App) ResponseHandler(hd ResponseHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hd(app.db, w, r)
	}
}

func (app *App) Get(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	app.routes.HandleFunc(path, fn).Methods("GET")
}

func (app *App) Post(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	app.routes.HandleFunc(path, fn).Methods("POST")
}

func (app *App) Put(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	app.routes.HandleFunc(path, fn).Methods("PUT")
}

func (app *App) Delete(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	app.routes.HandleFunc(path, fn).Methods("DELETE")
}
