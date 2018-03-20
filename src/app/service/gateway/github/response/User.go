package response

import (
	"app/service/gateway/github/resource"
)

// User - Canonical response which contains all info from several github end-points.
type User struct {
	Profile resource.Profile `json:"profile"`
	Repos   []resource.Repo  `json:"repos"`
	Orgs    []resource.Org   `json:"orgs"`
}
