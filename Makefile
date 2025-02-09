root?=$(abspath .)
name?=$(notdir $(root))
srcDir?=$(root)/cmd/$(name)
buildDir?=$(root)/bin
ldflags?="-s -w -X main.version=`git tag | tail -n1`"
prefix?=~/go/bin
internalFlags?="-s -w -X main"
default: test

all: build install

fetch:
	cd $(root)/internal; go run ./fetch
	# go run ./internal/fetch
	# go run $(root)/internal/fetch
	# cd $(root)

generate:
	go generate ./...

test: generate
	go test -v ./...

build:
	go build -o $(buildDir)/$(name) -ldflags=$(ldflags) $(srcDir)

install: build
	ln -fs $(buildDir)/$(name) $(prefix)/$(name)
