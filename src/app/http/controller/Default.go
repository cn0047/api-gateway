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
	message := protocol.Error(501, "Not Implemented.")
	protocol.Response(w, message)
}
