package main

import (
	"net/http"
)

var users = map[string]*User{}

type User struct {
	id   string "id"
	name string "name"
}

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// rw.Write([]byte("hello"))

		switch r.Method {
		case http.MethodGet:

			// json.NewEncod zer(rw).Encode()

		}

	})
	http.ListenAndServe(":8080", nil)

}
