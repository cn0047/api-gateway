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
