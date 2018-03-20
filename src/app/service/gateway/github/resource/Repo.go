package resource

// Repo - Structure which represents github repository.
type Repo struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}
