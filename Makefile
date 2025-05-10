.PHONY: test build tidy

build:
	go build -o bin/siwx-go .

test:
	go test ./...

tidy:
	go mod tidy
