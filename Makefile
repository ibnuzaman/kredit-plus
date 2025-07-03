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

## dev/install: install air for reloading on file changes and golangci-lint for linting
.PHONY: dev/install
dev/install:
	go install github.com/air-verse/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

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

## docs: generate swagger documentation
.PHONY: docs
docs:
	swag init .
ifeq ($(shell uname), Darwin)
	gsed -i "s/x-nullable/nullable/g" ./docs/docs.go
	gsed -i "s/x-omitempty/omitempty/g" ./docs/docs.go
	gsed -i "s/x-nullable/nullable/g" ./docs/swagger.json
	gsed -i "s/x-omitempty/omitempty/g" ./docs/swagger.json
	gsed -i "s/x-nullable/nullable/g" ./docs/swagger.yaml
	gsed -i "s/x-omitempty/omitempty/g" ./docs/swagger.yaml
else
	sed -i "s/x-nullable/nullable/g" ./docs/docs.go
	sed -i "s/x-omitempty/omitempty/g" ./docs/docs.go
	sed -i "s/x-nullable/nullable/g" ./docs/swagger.json
	sed -i "s/x-omitempty/omitempty/g" ./docs/swagger.json
	sed -i "s/x-nullable/nullable/g" ./docs/swagger.yaml
	sed -i "s/x-omitempty/omitempty/g" ./docs/swagger.yaml
endif
