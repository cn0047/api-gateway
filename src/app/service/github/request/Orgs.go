package request

import (
	"app/service/github/resource"
)

// GetOrgsJob describes job to get information about orgs.
func GetOrgsJob(userName string) Job {
	return Job{
		Name:     "orgs",
		Path:     "/orgs",
		Data:     &[]resource.Org{},
		UserName: userName,
	}
}
