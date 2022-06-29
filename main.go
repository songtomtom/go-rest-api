package main

import (
	"net/http"

	"songtomtom/rest_api/app"
)

func main() {
	app := app.App{}
	app.Init()
	http.ListenAndServe(":8080", nil)
}
