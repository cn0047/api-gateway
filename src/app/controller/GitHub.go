package controller

import (
	"net/http"

	"app/protocol"
	"app/service/github"
)

// GitHub controller, which contains all stuff related to github end-point.
type GitHub struct {
}

func (u GitHub) registerRoutes() {
	http.HandleFunc("/github/users/", u.handleRequest)
}

func (u GitHub) handleRequest(w http.ResponseWriter, r *http.Request) {
	// Default HTTP message (Error 501).
	message := protocol.Error(501, "Not Implemented.")

	if r.Method == "GET" {
		userName := r.URL.Path[len("/github/users/"):]
		data, err := github.GetUserInfo(userName)
		if len(err) == 0 {
			message = protocol.Success(200, data)
		} else {
			message = protocol.Exception(err)
		}
	}

	protocol.Response(w, message)
}
