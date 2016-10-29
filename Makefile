# Makefile - Ancestral friend, new language, same purpose
# When in doubt just `make`
# Some targets belong to:
# 	- GoTool Hepers - https://ariejan.net/2015/10/03/a-makefile-for-golang-cli-tools/
# 	- Goa Regen Helpers - https://github.com/goadesign/examples/tree/master/code_regen

BINARY=copper-mantis
VERSION=0.1.0

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -type f -name '*.go' ! -path '$(SOURCEDIR)/design/*'  ! -path '$(SOURCEDIR)/vendor/*' )
BUILD_TIME=$(shell date +%FT%T%z)
GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GO_NAMESPACE=github.com/CopperMantis/CopperMantis
GO_LDFLAGS= -X $(GO_NAMESPACE)/main.Version=$(VERSION)
GO_LDFLAGS+= -X $(GO_NAMESPACE)/main.BuildTime=$(BUILD_TIME)
GO_LDFLAGS+= -X $(GO_NAMESPACE)/main.Commit=$(GIT_COMMIT)

# Build
$(BINARY): vendor main.go $(SOURCES)
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "$(GO_LDFLAGS)" -o $(BINARY)

.PHONY: vendor clean clean-vendor
vendor:
	glide install

clean:
	@[ -f $(BINARY) ] && rm $(BINARY) || true

clean-vendor:
	@[ -d vendor ] && rm -r vendor || true

# Test
.PHONY: test lint
test: vendor lint
	go test -v $(SOURCES) -cover

lint:
	go fmt $(SOURCES)
	go vet $(SOURCES)

# Goa realated and misc targets
include build/*.mk
