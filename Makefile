start: build
	@ ./bin/main

build:
	@go build -o ./bin ./cmd/api/main.go

