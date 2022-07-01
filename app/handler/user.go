package handler

import (
	"encoding/json"
	"net/http"
	"songtomtom/rest_api/app/model"

	"gorm.io/gorm"
)

func FindAll(db *gorm.DB, rw http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)

	res, err := json.Marshal(users)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(res))
}
