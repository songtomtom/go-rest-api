package handler

import (
	"encoding/json"
	"net/http"
)

func index() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello"))
	}
}

func responseSuccess(payload interface{}) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(payload)
		if err != nil {
			responseError(err)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(res))
	}
}

func responseError(err error) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}
}
