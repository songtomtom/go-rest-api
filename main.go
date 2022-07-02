package main

import "songtomtom/rest_api/app"

func main() {
	app := &app.App{}
	app.Run(8080)
}
