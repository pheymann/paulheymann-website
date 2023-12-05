# Contract Testing Example Project
This is an example of how to use Contract Testing to cover the frontend (React) and backend (Golang).

## How To
### Set it up
Install the server:

```shell
go mod vendor
```

Install the app:

```shell
cd app && npm install
```

### Run it locally
Start the server:

```shell
go run cmd/runmovieserver/main.go
```

Start the app:

```shell
cd app && npm start
```

### Run CDC tests
For the server:

```shell
go test ./...
```

For the app:

```shell
cd app && npm test
```

## Explanation
### Repo Structure
This is a monorepo. Server and app code live in the same place. In the root, you will find the standard Golang project setup (`cmd`, `internal`, etc.). Besides that, there is also an `app` folder storing all the React/JS code for the app.

Most importantly, contract tests are defined as `YAML` and are located under `contracts/`.
