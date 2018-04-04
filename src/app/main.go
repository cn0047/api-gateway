package main

import (
	"net/http"

	"app/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8080", nil)
}
