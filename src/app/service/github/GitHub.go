package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"app/service/github/resource"
)

const (
	gitHubAPI = "https://api.github.com/users/"
	timeout   = 10000 // milliseconds
)

var (
	httpClient = http.Client{Timeout: time.Millisecond * timeout}
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
			err := takeData(endPoint, &data)
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

// takeData gets data from github API.
func takeData(URI string, p interface{}) error {
	resp, err := httpClient.Get(URI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return err
	}

	return nil
}
