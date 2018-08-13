package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	gitHubAPI = "https://api.github.com/users/"
	timeout   = 10000 // milliseconds
)

var (
	httpClient = http.Client{Timeout: time.Millisecond * timeout}
	routes     = map[string]string{
		"profile": "",
		"repos":   "/repos",
		"orgs":    "/orgs",
	}
)

// GetURI func to get URI for particular username and route.
func GetURI(userName string, route string) string {
	suffix := routes[route]
	return gitHubAPI + userName + suffix
}

// UnmarshalData gets data from github API.
func UnmarshalData(URI string, p *interface{}) error {
	resp, err := httpClient.Get(URI)
	if err != nil {
		return fmt.Errorf("failed to perform request, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		f := "received response code: %v, but expected: %v, status: %v"
		return fmt.Errorf(f, resp.StatusCode, http.StatusOK, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body, error: %v", err)
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return fmt.Errorf("failed to perform json.Unmarshal, error: %v", err)
	}

	return nil
}
