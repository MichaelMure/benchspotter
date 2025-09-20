TAG:=$(shell git name-rev --name-only --tags HEAD)
LDFLAGS:=-X benchspotter/commands.version="${TAG}"

all: build

.PHONY: build
build:
	go generate
	go build -ldflags "$(LDFLAGS)" .

.PHONY: test
test:
	go test -v -race=. ./...