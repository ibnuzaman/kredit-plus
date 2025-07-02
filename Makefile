## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	go build -o /tmp/${binary_name} .

## build/run: build and run the application
.PHONY: build/run
build/run: build run

## run: run the application
.PHONY: run
run:
	export ENV=development && /tmp/${binary_name}

## dev: run the application with reloading on file changes
.PHONY: dev
dev:
	export ENV=development && air -c .air.toml

## dev/run: run the docs, lint, and build run
.PHONY: dev/run
dev/run: docs lint build run

## lint: run linters on the code
.PHONY: lint
lint:
	go fmt ./...
	go vet ./...
	swag fmt -d .
	golangci-lint run -v --fast

## clean: clean up the project
.PHONY: clean
clean:
	rm -f /tmp/${binary_name}
	rm -f /tmp/coverage.out
