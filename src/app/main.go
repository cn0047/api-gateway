package main

import (
	"net/http"

	"app/controller"
)

var (
	defaultController controller.Default
	gitHubController  controller.GitHub
)

func main() {
	defaultController.RegisterRoutes()
	gitHubController.RegisterRoutes()

	http.ListenAndServe(":8080", nil)
}
