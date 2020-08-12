BIN      = $(CURDIR)/bin
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)


.PHONY: all

all:
	go build -tags release \
	          -ldflags "-X 'github.com/gruffwizard/kabnet/cmd.Version=$(VERSION)'" \
						-o $(BIN)/kabnet 

compile:
	go build -o bin/main main.go
test:
	go test ./...
run:
	go run main.go

package:
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
