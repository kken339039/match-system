.PHONY: start_dev start_prod format lint build

format:
	@gofmt -e -s -w -l ./

lint:
	@golangci-lint run -v ./... --timeout 3m0s

start_dev:
	@air

build:
	@go build -o build/server cmd/main.go

start_prod: build
	@build/api .env