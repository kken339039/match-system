.PHONY: start_dev format lint build mocks test

format:
	@gofmt -e -s -w -l ./

lint:
	@golangci-lint run -v ./... --timeout 3m0s

start_dev:
	@air

build:
	@go build -o build/server cmd/main.go

test:
	@rm -rf ./coverage && mkdir coverage
	@PROJECT_ROOT=$(PWD) ENVIRONMENT=test go test -parallel 4 -race -tags $(BUILD_TAGS_API) -covermode=atomic -coverprofile=coverage/coverage.out -coverpkg=./internal/... ./tests/...

mocks:
	go mod tidy
	@docker run -v "$(PWD)":/src -w /src vektra/mockery --dir=interfaces/ --all