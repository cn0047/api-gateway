package protocol

// Message - Canonical HTTP protocol level message.
// Any HTTP response must be provided using this message.
type Message struct {
	Error   ErrorMessage   `json:"error"`
	Success SuccessMessage `json:"success"`
}

// GetError - get error message structure.
func (r Message) GetError() ErrorMessage {
	return r.Error
}

// GetSuccess - get success message structure.
func (r Message) GetSuccess() SuccessMessage {
	return r.Success
}
