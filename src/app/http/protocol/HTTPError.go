package protocol

// HTTPErrorMessage - Canonical HTTP error message.
// This structure must be used to handle any error.
type HTTPErrorMessage struct {
	Code        int         `json:"code"`
	Description interface{} `json:"description"`
	Data        interface{} `json:"data"`
}

// HTTPError - Wrapper for HTTP error message,
// which helps to handle errors in more elegant way.
func HTTPError(code int, desc string) HTTPMessage {
	e := HTTPErrorMessage{Code: code, Description: desc}

	return HTTPMessage{Error: e, Success: HTTPSuccessMessage{}}
}

// HTTPException - Wrapper for errors gained from "panic".
// In addition to regular error information,
// this one provides additional custom payload.
func HTTPException(data interface{}) HTTPMessage {
	e := HTTPErrorMessage{
		Code:        500,
		Description: "Internal Server Error.",
		Data:        data,
	}

	return HTTPMessage{Error: e, Success: HTTPSuccessMessage{}}
}
