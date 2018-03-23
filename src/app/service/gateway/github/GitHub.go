package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"app/service/gateway/github/resource"
	"app/service/gateway/github/response"
)

const (
	gitHubAPIURL                = "https://api.github.com/users/"
	gitHubAPITimeoutMillisecond = 10000
)

// GetUserInfo - Get from github information about user, user's repos and organizations.
func GetUserInfo(userName string) (response.User, []string) {
	var user response.User
	var errorsList []string
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		data, err := getProfileInfo(userName)
		if err == nil {
			user.Profile = data
		} else {
			errorsList = append(errorsList, err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		data, err := getReposInfo(userName)
		if err == nil {
			user.Repos = data
		} else {
			errorsList = append(errorsList, err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		data, err := getOrgsInfo(userName)
		if err == nil {
			user.Orgs = data
		} else {
			errorsList = append(errorsList, err.Error())
		}
	}()

	wg.Wait()

	return user, errorsList
}

// Perform request into github API.
// This method only performs request and returns response body which intended to contain payload,
// hence it's possible to use this method for any github API end-point.
func getDataFromGitHub(endPoint string) ([]byte, error) {
	client := http.Client{Timeout: time.Millisecond * gitHubAPITimeoutMillisecond}

	resp, err := client.Get(gitHubAPIURL + endPoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("resource not found")
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Get info about user (general profile information)
func getProfileInfo(userName string) (resource.Profile, error) {
	var profile resource.Profile
	rawRepos, err := getDataFromGitHub(userName)
	if err != nil {
		return profile, err
	}
	json.Unmarshal(rawRepos, &profile)

	return profile, nil
}

// Get info about user's repositories.
func getReposInfo(userName string) ([]resource.Repo, error) {
	var repos []resource.Repo
	rawRepos, err := getDataFromGitHub(userName + "/repos")
	if err != nil {
		return repos, err
	}
	json.Unmarshal(rawRepos, &repos)

	return repos, nil
}

// Get info about user's organizations.
func getOrgsInfo(userName string) ([]resource.Org, error) {
	var orgs []resource.Org
	rawOrgs, err := getDataFromGitHub(userName + "/orgs")
	if err != nil {
		return orgs, err
	}
	json.Unmarshal(rawOrgs, &orgs)

	return orgs, nil
}
