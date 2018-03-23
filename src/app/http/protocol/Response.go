package protocol

import (
	"encoding/json"
	"net/http"
)

// Response - Canonical way to sends HTTP message response to client.
func Response(w http.ResponseWriter, response Message) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(getHTTPCode(response))

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("RUNTIME-ERROR-JSON-1: " + err.Error())
	}
}

// getHTTPCode - get HTTP status code from provided HTTPMessage.
func getHTTPCode(response Message) int {
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
