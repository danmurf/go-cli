.DEFAULT_GOAL := build
COMMAND_BUILD_FILE := shout.go
COMMAND_NAME := shout

build: lint test
	go build -o bin/${COMMAND_NAME} cmd/${COMMAND_BUILD_FILE}

lint:
	go fmt github.com/danmurf...

test:
	go test ./...
