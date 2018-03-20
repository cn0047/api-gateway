package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"app/service/gateway/github/response"
	"app/service/gateway/github/resource"
)

const (
	gitHubApiURL = "https://api.github.com/users/"
)

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

func req(endPoint string) []byte {
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

func getProfileInfo(userName string) resource.Profile {
	var profile resource.Profile
	rawRepos := req(userName)
	json.Unmarshal(rawRepos, &profile)

	return profile
}

func getReposInfo(userName string) []resource.Repo {
	var repos []resource.Repo
	rawRepos := req(userName + "/repos")
	json.Unmarshal(rawRepos, &repos)

	return repos
}

func getOrgsInfo(userName string) []resource.Org {
	var orgs []resource.Org
	rawOrgs := req(userName + "/orgs")
	json.Unmarshal(rawOrgs, &orgs)

	return orgs
}
