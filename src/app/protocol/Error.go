package protocol

// ErrorMessage - Canonical HTTP error message.
// This structure must be used to handle any error.
type ErrorMessage struct {
	Code        int         `json:"code"`
	Description interface{} `json:"description"`
	Data        interface{} `json:"data"`
}

// Error - Wrapper for HTTP error message,
// which helps to handle errors in more elegant way.
func Error(code int, desc string) Message {
	e := ErrorMessage{Code: code, Description: desc}

	return Message{Error: e, Success: SuccessMessage{}}
}

// Exception - Wrapper for errors gained from "panic".
// In addition to regular error information,
// this one provides additional custom payload.
func Exception(data interface{}) Message {
	e := ErrorMessage{
		Code:        500,
		Description: "Internal Server Error.",
		Data:        data,
	}

	return Message{Error: e, Success: SuccessMessage{}}
}
