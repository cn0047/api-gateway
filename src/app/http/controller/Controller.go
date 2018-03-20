package controller

var (
	defaultController Default
	gitHubController  GitHub
)

// Startup - Init all routes for all controllers.
func Startup() {
	defaultController.registerRoutes()
	gitHubController.registerRoutes()
}
