package request

import (
	"app/service/github/resource"
)

// GetReposJob describes job to get information about repos.
func GetReposJob(userName string) Job {
	return Job{
		Name:     "repos",
		Path:     "/repos",
		Data:     &[]resource.Repo{},
		UserName: userName,
	}
}
