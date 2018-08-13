API Gateway
-

[![CircleCI](https://circleci.com/gh/cn007b/api-gateway.svg?style=svg)](https://circleci.com/gh/cn007b/api-gateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/cn007b/api-gateway)](https://goreportcard.com/report/github.com/cn007b/api-gateway)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a75fd2bff6ae4365b6535126ff429621)](https://www.codacy.com/app/cn007b/api-gateway?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cn007b/api-gateway&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/89a5bcee77752ad8a6ec/maintainability)](https://codeclimate.com/github/cn007b/api-gateway/maintainability)

This is super simple `API Gateway` for `github`
<br>which helps receive data from several github API endpoints in one call.

## Usage

Run in shell next command:

````bash
go run src/app/main.go
# or
go run -race src/app/main.go
````

Run query:

````bash
curl 'http://localhost:8080/github/users/cn007b'
````

As result you'll see something like this:

````json
{
  "error": null,
  "success": {
    "code": 200,
    "data": {
      "orgs": {
        "data": [
          {
            "login": "thepkg",
            "description": "Go Packages."
          }
        ],
        "error": null
      },
      "profile": {
        "data": {
          "login": "cn007b",
          "name": "Vladimir Kovpak",
          "location": "Kyiv, Ukraine"
        },
        "error": null
      },
      "repos": {
        "data": [
          {
            "name": "api-gateway",
            "full_name": "cn007b/api-gateway"
          }
        ],
        "error": null
      }
    }
  }
}
````
