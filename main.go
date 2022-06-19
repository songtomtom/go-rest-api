package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	})

	http.HandleFunc("/bar", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello Bar")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
