package github

import (
	"sync"

	"app/service/gateway/externalapi"
	"app/service/gateway/github/resource"
)

const (
	gitHubAPI = "https://api.github.com/users/"
)

// GetUserInfo - Get from github information about user, user's repos and organizations.
func GetUserInfo(userName string) (map[string]interface{}, map[string]interface{}) {
	// Errors map.
	e := make(map[string]interface{})

	// Result payload map. Contains structures which will be populated by data from github.
	payload := make(map[string]interface{})
	payload["profile"] = &resource.Profile{}
	payload["repos"] = &[]resource.Repo{}
	payload["orgs"] = &[]resource.Org{}

	// Map: payload key to end-point suffix.
	m := make(map[string]string)
	m["profile"] = ""
	m["repos"] = "/repos"
	m["orgs"] = "/orgs"

	var wg sync.WaitGroup
	wg.Add(len(payload))

	for k, suffix := range m {
		key := k
		endPoint := gitHubAPI + userName + suffix
		data := payload[key]
		go func() {
			defer wg.Done()
			err := externalapi.TakeData(endPoint, &data)
			if err == nil {
				payload[key] = data
			} else {
				e[key] = err.Error()
			}
		}()
	}

	wg.Wait()

	return payload, e
}
