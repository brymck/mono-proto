PROJECT_NAME := mono-proto
GO_FILES := main.go $(shell find internal pkg -name "*.go")

all: $(PROJECT_NAME)

$(PROJECT_NAME):
	go build ./...

profile.out: $(GO_FILES)
	go test -coverprofile=profile.out ./...

test: profile.out

cover: test
	go tool cover -html=profile.out

clean:
	rm mono-proto profile.out

.PHONY: all test cover clean
