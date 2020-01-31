PROJECT_NAME := mono-proto
GO_FILES := main.go $(shell find internal pkg -name "*.go")

all: $(PROJECT_NAME)

$(PROJECT_NAME):
	go build ./...

coverage.txt: $(GO_FILES)
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

test: coverage.txt

cover: test
	go tool cover -html=coverage.txt

clean:
	rm mono-proto coverage.txt

.PHONY: all test cover clean
