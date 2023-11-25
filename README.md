# Match Sysyem

This is a matching system implemented in Golang, with an HTTP server providing API calls. The system allows users to add, remove, and query possible matches for both males and females based on specified criteria.

## Features

- Add a new user to the matching system and find possible matches.
- Remove a user from the matching system.
- Query for the most N possible matched single people.

## Matching Rules

- A single person has four input parameters: name, height, gender, and number of wanted dates.
- Boys can only match girls who have a lower height. Conversely, girls match boys who are taller.
- Once the girl and boy match, they both use up one date. When their number of dates becomes zero, they are removed from the matching system.

## Get Started

1. Clone repo

```
git clone https://github.com/kken339039/match-system.git
cd matching-system
```

2. Build and run the Docker container:

```
docker build -t matching-system .
docker run -p 3000:3000 matching-system
```


### Install required packages

```
brew install golangci-lint
go install github.com/cosmtrek/air@latest
```

### Setup the HTTP server in local

1. install Go 1.20.2
2. copy `.env.example` as `.env` and update the values
3. Run `go mod download` to download the go modules
4. go install github.com/cosmtrek/air@latest
5. Run `make start_dev` to start the HTTP server
6. Use postman or to call api and start development

## Project Structure

The project follows a structured layout for better organization. Below is an overview of the project structure:

```
/match-system
  /cmd
    main.go
  /interfaces
  /internal
    /store
    /user
  /mocks
  /plugins
    /env
    /logger
    /http_server
  /tests
  go.mod
  go.sum
```

### main.go
Main application entry point

### interfaces
Defined modesl, service, store functional

### internal
Implement application match male/female, remove, query logic
- store: In memory data struct
- user: Implement User match logic and controller

### mocks
- Mock interfaces for unit testing

### plugins
Project plugins and extensions
- env: Application environment variables
- logger: Application logging
- http_server: Implementent HTTP server

### tests
Unit Test

## Format & Lint Code

To format & lint code follow golangci.yml, then you can fix it if any warning or error. use the following command:

```
make format
make lint
```

## Unit Tests

If you test any new test and need mock service or store some func, you can run the following command to generate mock interface for mock test return.

```
make mocks
```

To run unit tests, use the following command:

```
make test
``````
