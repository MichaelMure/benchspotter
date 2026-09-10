TAG:=$(shell git name-rev --name-only --tags HEAD)
LDFLAGS:=-X github.com/MichaelMure/benchspotter/commands.version="${TAG}"

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
	vhs doc/demo/tapes/bench.tape
	vhs doc/demo/tapes/compare.tape
	vhs doc/demo/tapes/show-diff.tape
	vhs doc/demo/tapes/show-cpu.tape
	vhs doc/demo/tapes/show-escape.tape
	vhs doc/demo/tapes/trend.tape
	vhs doc/demo/tapes/optimize.tape