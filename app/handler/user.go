package handler

import (
	"encoding/json"
	"net/http"
	"songtomtom/rest_api/app/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FindAll(db *gorm.DB, w http.ResponseWriter, r *http.Request) func(http.ResponseWriter, *http.Request) {
	users := []model.User{}
	db.Find(&users)
	return responseSuccess(users)
}

func FindById(db *gorm.DB, w http.ResponseWriter, r *http.Request) func(http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)

	user := model.User{}
	if err := db.First(&user, "id = ?", vars["id"]); err != nil {
		return responseError(err.Error)
	}

	return responseSuccess(user)

}

func Create(db *gorm.DB, w http.ResponseWriter, r *http.Request) func(http.ResponseWriter, *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		return responseError(err)
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err == nil {
		return responseError(err)
	}

	return responseSuccess(user)

}

func Update(db *gorm.DB, w http.ResponseWriter, r *http.Request) func(http.ResponseWriter, *http.Request) {

	vars := mux.Vars(r)

	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		return responseError(err)
	}
	defer r.Body.Close()

	if err := db.First(&user, "id = ?", vars["id"]); err != nil {
		return responseError(err.Error)
	}

	if err := db.Save(&user).Error; err == nil {
		return responseError(err)
	}

	return responseSuccess(user)

}

func Delete(db *gorm.DB, w http.ResponseWriter, r *http.Request) func(http.ResponseWriter, *http.Request) {

	vars := mux.Vars(r)

	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		return responseError(err)
	}
	defer r.Body.Close()

	if err := db.First(&user, "id = ?", vars["id"]); err != nil {
		return responseError(err.Error)
	}

	if err := db.Delete(&user).Error; err == nil {
		return responseError(err)
	}

	return responseSuccess(user)

}
