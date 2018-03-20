package protocol

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse - Canonical way to sends HTTP message response to client.
func HTTPResponse(w http.ResponseWriter, response HTTPMessage) {
	w.WriteHeader(getHTTPCode(response))

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("RUNTIME-ERROR-JSON-1: " + err.Error())
	}
}

// Gets HTTP status code from provided HTTPMessage.
func getHTTPCode(response HTTPMessage) int {
	var httpCode int

	errorCode := response.GetError().Code
	successCode := response.GetSuccess().Code
	if successCode > 0 {
		httpCode = successCode
	}
	// This block was written in this way intentionally,
	// because error code has higher priority and may overlap success code.
	if errorCode > 0 {
		httpCode = errorCode
	}

	return httpCode
}