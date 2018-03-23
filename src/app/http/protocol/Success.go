package protocol

// SuccessMessage - Canonical HTTP success message.
// This structure must be used to reply with success message.
// This structure is - container for payload shipping.
type SuccessMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Success - Wrapper for HTTP success message,
// which helps to provide payload in more elegant way.
func Success(code int, data interface{}) Message {
	s := SuccessMessage{Code: code, Data: data}

	return Message{Error: ErrorMessage{}, Success: s}
}
