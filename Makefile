.DEFAULT_GOAL := build

OS := $(shell uname -s | tr A-Z a-z)

build:
	mkdir -p ./bin
	go build -o ./bin/oeis .

tests:
	go test ./...