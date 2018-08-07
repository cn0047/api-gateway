package controller

import (
	"net/http"
	"github.com/thepkg/rest"
)

// Default controller, which works as fallback end-point.
type Default struct {
}

func (d Default) registerRoutes() {
	rest.GET("/", d.handleRequest)
}

func (d Default) handleRequest(w http.ResponseWriter, r *http.Request) {
	rest.Error(w, http.StatusNotImplemented, "Not Implemented.")
}
