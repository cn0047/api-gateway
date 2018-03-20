package main

import (
	"net/http"

	"app/http/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8080", nil)
}
