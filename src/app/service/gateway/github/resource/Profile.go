package resource

// Profile - Structure which represents github profile (general information).
type Profile struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
