API Gateway
-

This is super simple `API Gateway` for `github`
<br>which helps receive data from several github API endpoints in one call.

## Usage

Run `docker` container:

````
docker run -it --rm -p 8080:8080 -v $PWD:/app -w /app -e GOPATH='/app' golang:latest go run src/app/main.go
````

Run query:

````
curl 'http://localhost:8080/github/users/cn007b'
````

As result you'll see something like this:

````
{
  "error": {
    "code": 0,
    "description": null,
    "data": null
  },
  "success": {
    "code": 200,
    "data": {
      "profile": {
        "login": "cn007b",
        "name": "Vladimir Kovpak",
        "location": "Kiev, Ukraine"
      },
      "repos": [
        {
          "name": "api-gateway",
          "full_name": "cn007b/api-gateway"
        },
        {
          "name": "benchmark-postgres-mongo",
          "full_name": "cn007b/benchmark-postgres-mongo"
        }
      ],
      "orgs": [
        {
          "login": "thisiskint",
          "description": ""
        }
      ]
    }
  }
}
````
