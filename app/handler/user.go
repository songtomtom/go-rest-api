package handler

import (
	"net/http"
	"songtomtom/rest_api/app/model"

	"gorm.io/gorm"
)

func FindAll(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	users := []model.User{}
	db.Find(&users)
	return responseSuccess(users)
}

func FindById(db *gorm.DB) func(http.ResponseWriter, *http.Request) {

	user := model.User{}
	err := db.First(&user)
	if err != nil {
		return responseError(err.Error)
	}

	return responseSuccess(user)

}

func Create(db *gorm.DB) func(http.ResponseWriter, *http.Request) {

	user := model.User{}

	if err != nil {
		return responseError(err.Error)
	}

	return responseSuccess(user)

}

func Update(db *gorm.DB) func(http.ResponseWriter, *http.Request) {

	user := model.User{}
	err := db.First(&user)
	if err != nil {
		return responseError(err.Error)
	}

	return responseSuccess(user)

}

func Delete(db *gorm.DB) func(http.ResponseWriter, *http.Request) {

	user := []model.User{}
	err := db.First(&user)
	if err != nil {
		return responseError(err.Error)
	}

	return responseSuccess(user)

}
