include ./globals.mk

.PHONY: test
test: ## Tests the lib
	cd src
	go test ./...

build: test ## Builds the artifacts
	mkdir -p build
	cd src
	go build .
	mv markov ../build