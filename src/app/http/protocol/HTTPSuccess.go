package protocol

// HTTPSuccessMessage - Canonical HTTP success message.
// This structure must be used to reply with success message.
// This structure is - container for payload shipping.
type HTTPSuccessMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// HTTPSuccess - Wrapper for HTTP success message,
// which helps to provide payload in more elegant way.
func HTTPSuccess(code int, data interface{}) HTTPMessage {
	s := HTTPSuccessMessage{Code: code, Data: data}

	return HTTPMessage{Error: HTTPErrorMessage{}, Success: s}
}
