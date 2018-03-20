package controller

var (
	defaultController Default
	gitHubController GitHub
)

func Startup() {
	defaultController.registerRoutes()
	gitHubController.registerRoutes()
}
