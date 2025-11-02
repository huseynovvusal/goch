BINARY_NAME=goch

build:
	go build -o bin/$(BINARY_NAME) main.go

test:
	go test ./...

.PHONY: build test