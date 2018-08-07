package controller

import (
	"net/http"
	"github.com/thepkg/rest"

	"app/service/github"
)

// GitHub controller, which contains all stuff related to github end-point.
type GitHub struct {
}

func (u GitHub) registerRoutes() {
	rest.GET("/github/users/", u.handleRequest)
}

func (u GitHub) handleRequest(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Path[len("/github/users/"):]
	data, err := github.GetUserInfo(userName)
	if len(err) == 0 {
		rest.Success(w,http.StatusOK, data)
	} else {
		rest.Error(w, http.StatusBadRequest, err)
	}
}
