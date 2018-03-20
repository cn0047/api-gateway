package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"app/service/gateway/github/resource"
	"app/service/gateway/github/response"
)

const (
	gitHubApiURL = "https://api.github.com/users/"
);

// Get from github information about user, user's repos and organizations.
func GetUserInfo(userName string) response.User {
	var user response.User
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		user.Profile = getProfileInfo(userName)
	}()
	go func() {
		defer wg.Done()
		user.Repos = getReposInfo(userName)
	}()
	go func() {
		defer wg.Done()
		user.Orgs = getOrgsInfo(userName)
	}()

	wg.Wait()

	return user
}

// Perform request into github API.
// This method only performs request and returns response body which intended to contain payload,
// hence it's possible to use this method for any github API end-point.
func getDataFromGitHub(endPoint string) []byte {
	resp, err := http.Get(gitHubApiURL + endPoint)
	if err != nil {
		panic("RUNTIME-ERROR-GITHUB-API-1: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("RUNTIME-ERROR-GITHUB-API-2: " + err.Error())
	}

	return body
}

// Get info about user (general profile information)
func getProfileInfo(userName string) resource.Profile {
	var profile resource.Profile
	rawRepos := getDataFromGitHub(userName)
	json.Unmarshal(rawRepos, &profile)

	return profile
}

// Get info about user's repositories.
func getReposInfo(userName string) []resource.Repo {
	var repos []resource.Repo
	rawRepos := getDataFromGitHub(userName + "/repos")
	json.Unmarshal(rawRepos, &repos)

	return repos
}

// Get info about user's organizations.
func getOrgsInfo(userName string) []resource.Org {
	var orgs []resource.Org
	rawOrgs := getDataFromGitHub(userName + "/orgs")
	json.Unmarshal(rawOrgs, &orgs)

	return orgs
}
