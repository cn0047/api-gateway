package controller

import (
	"github.com/thepkg/rest"
	"net/http"
)

// Default controller, which works as fallback end-point.
type Default struct {
}

// RegisterRoutes registers HTTP routes handlers.
func (d Default) RegisterRoutes() {
	rest.GET("/", d.handleRequest)
}

func (d Default) handleRequest(w http.ResponseWriter, r *http.Request) {
	rest.Error(w, http.StatusNotImplemented, "Not Implemented.")
}
