package request

import (
	"app/service/github/resource"
)

// GetProfileJob describes job to get information about user profile.
func GetProfileJob(userName string) Job {
	return Job{
		Name:     "profile",
		Path:     "/",
		Data:     &resource.Profile{},
		UserName: userName,
	}
}
