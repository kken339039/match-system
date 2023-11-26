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

## API documentation

you can generate doc, use the following command:

```
make swag
```

after run command, you can review on browser: `http://localhost:3000/swagger/index.html`

## System Design Documentation

This Match System is designed to efficiently match users based on their preferences, considering as gender, height, and count of wanted dates. The system is implemented in Golang and follows as microservice.

### Microservices:

- Users Service: Manages users relations, like adding users, removing users, and querying single users.

### Data Storage:

- User data is stored in-memory using a store(memory) structure. Each user contains information such as ID, name, height, gender, wanted dates, and a list of matched users.

### Matching Algorithm:

- Users are matched on gender and height preferences.
- When a new user added, the system found existing users to match suitable user.

### Documentation:

- API documentation is generated using Swagger for reading.

### Testing:

Unit tests are implemented to ensure the correctness of individual components, including controllers and services.

### Time Complexity of API Operations
#### AddUserAndMatch:

- Time Complexity: O(n), where n is the number of existing users.
- Found existing users to find suitable user based gender and height.

#### RemoveTargetUser:

- Time Complexity: O(n), where n is the number of existing users.
- Search the target user in the list of users and removing them.

#### QuerySingleUsers:

- Time Complexity: O(n), where n is the number of existing users.
- Searching for single users based on count of wanted date which is more than zero.

## TBD
### Data Store
- Select a persistent data source, like MySQL, PostgreSQL

### Security Measures
- Implement auth logic to secure API endpoints.

### Monitoring and Logging:
- Implement monitor and logging solutions to track system error and diagnose issues.

### CI/CD Pipeline:
- Set up a continuous integration/continuous deployment (CI/CD) pipeline for automated testing and deployment.

### Dependency Injection (DI) Support:
- Consider a Dependency Injection framework to achieve better code structuring, testability, and maintainability, like(uber/fx).
- DI helps reduce coupling between components, enhances code testability, and makes the code more extensible and modifiable.
