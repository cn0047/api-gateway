package response

import (
	"app/service/gateway/github/resource"
)

type User struct {
	Profile resource.Profile `json:"profile"`
	Repos []resource.Repo `json:"repos"`
	Orgs []resource.Org `json:"orgs"`
}
