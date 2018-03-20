package controller

import (
	"net/http"

	"app/http/protocol"
	"app/service/gateway/github"
)

// GitHub controller, which contains all stuff related to github end-point.
type GitHub struct {
}

func (u GitHub) registerRoutes() {
	http.HandleFunc("/github/users/", u.handleRequest)
}

func (u GitHub) handleRequest(w http.ResponseWriter, r *http.Request) {
	// Default HTTP message (Error 501).
	message := protocol.HTTPError(501, "Not Implemented.")

	defer func() {
		protocol.HTTPResponse(w, message)
	}()

	// Regardless panic reply to client with correct HTTP message.
	defer func() {
		if err := recover(); err != nil {
			message = protocol.HTTPException(err)
		}
	}()

	if r.Method == "GET" {
		userName := r.URL.Path[len("/github/users/"):]
		message = protocol.HTTPSuccess(200, github.GetUserInfo(userName))
	}
}
