.PHONY: run build deps

run:
	go run cmd/main.go

build:
	go build -o bin/testEx cmd/main.go

deps:
	go mod download