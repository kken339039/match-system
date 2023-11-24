.PHONY: start_dev format lint build

format:
	@gofmt -e -s -w -l ./

lint:
	@golangci-lint run -v ./... --timeout 3m0s

start_dev:
	@air

build:
	@go build -o build/server cmd/main.go
