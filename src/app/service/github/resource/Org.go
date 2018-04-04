package resource

// Org - Structure which represents github organization.
type Org struct {
	Login       string `json:"login"`
	Description string `json:"description"`
}
