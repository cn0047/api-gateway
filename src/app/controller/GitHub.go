package controller

import (
	"github.com/thepkg/rest"
	"net/http"

	"app/service/github"
)

// GitHub controller, which contains all stuff related to github end-point.
type GitHub struct {
}

// RegisterRoutes registers HTTP routes handlers.
func (u GitHub) RegisterRoutes() {
	rest.GET("/github/users/", u.handleRequest)
}

func (u GitHub) handleRequest(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Path[len("/github/users/"):]
	data := github.GetUserInfo(userName)
	rest.Success(w, http.StatusOK, data)
}
