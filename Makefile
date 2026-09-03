TAG:=$(shell git name-rev --name-only --tags HEAD)
LDFLAGS:=-X benchspotter/commands.version="${TAG}"

all: build

.PHONY: build
build:
	go generate
	go build -ldflags "$(LDFLAGS)" .

.PHONY: test
test:
	go test -v -race ./...

.PHONY: demo
demo:
	vhs demo/tapes/bench.tape
	vhs demo/tapes/compare.tape
	vhs demo/tapes/show-diff.tape
	vhs demo/tapes/show-cpu.tape
	vhs demo/tapes/show-escape.tape
	vhs demo/tapes/trend.tape
	vhs demo/tapes/optimize.tape