.DEFAULT_GOAL := build

OS := $(shell uname -s | tr A-Z a-z)


 ifeq (, $(shell which go))
 $(error "No go in $(PATH), consider checking out https://golang.org/doc/install for info")
 endif

build:
	mkdir -p ./bin
	go build -o ./bin/oeis .

test:
	go test ./...