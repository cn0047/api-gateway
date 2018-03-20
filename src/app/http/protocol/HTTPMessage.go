package protocol

// HTTPMessage - Canonical HTTP protocol level message.
// Any HTTP response must be provided using this message.
type HTTPMessage struct {
	Error   HTTPErrorMessage   `json:"error"`
	Success HTTPSuccessMessage `json:"success"`
}

// GetError - get error message structure.
func (r HTTPMessage) GetError() HTTPErrorMessage {
	return r.Error
}

// GetSuccess - get success message structure.
func (r HTTPMessage) GetSuccess() HTTPSuccessMessage {
	return r.Success
}
