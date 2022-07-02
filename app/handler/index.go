package handler

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func Index(_ *gorm.DB, _ http.ResponseWriter, _ *http.Request) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	}
}

func responseSuccess(payload interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(payload)
		if err != nil {
			responseError(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	}
}

func responseError(err error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
