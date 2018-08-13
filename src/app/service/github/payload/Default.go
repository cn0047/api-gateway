package payload

// Default describes default payload struct.
type Default struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}
