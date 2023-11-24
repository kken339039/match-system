.PHONY: start_dev format lint build mocks

format:
	@gofmt -e -s -w -l ./

lint:
	@golangci-lint run -v ./... --timeout 3m0s

start_dev:
	@air

build:
	@go build -o build/server cmd/main.go

mocks:
	go mod tidy
	@docker run -v "$(PWD)":/src -w /src vektra/mockery --dir=interfaces/ --all