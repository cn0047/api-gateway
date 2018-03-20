package controller

import (
	"net/http"

	"app/http/protocol"
)

// Default controller, which works as fallback end-point.
type Default struct {
}

func (d Default) registerRoutes() {
	http.HandleFunc("/", d.handleRequest)
}

func (d Default) handleRequest(w http.ResponseWriter, r *http.Request) {
	message := protocol.HTTPError(501, "Not Implemented.")
	protocol.HTTPResponse(w, message)
}
