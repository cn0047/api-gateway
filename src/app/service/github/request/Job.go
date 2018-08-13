package request

// Job describes job which gets data from github.
type Job struct {
	Name     string
	Path     string
	Data     interface{}
	UserName string
	Error    error
}
