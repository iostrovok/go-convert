CURDIR := $(shell pwd)
GOBIN := $(CURDIR)/bin/
ENV:=GOBIN=$(GOBIN)
DIR:=FILE_DIR=$(CURDIR)/testfiles TEST_SOURCE_PATH=$(CURDIR)
GODEBUG:=GODEBUG=gocacheverify=1

##
## List of commands:
##

## default:
all: mod test

deps:
	@echo "======================================================================"
	@echo 'MAKE: deps...'
	@mkdir -p $(GOBIN)
	@$(ENV) go get -u golang.org/x/lint/golint

test: tests

tests:
	@echo "Run test -race ./..."
	@$(DIR) $(GODEBUG) go test -race ./...

tests-cover:
	@echo "Run test  -cover -race -coverprofile=./coverage.out ./..."
	@$(DIR) $(GODEBUG) go test -cover -race -coverprofile=./coverage.out ./...
	go tool cover -html=./coverage.out -o ./coverage.html
	rm ./coverage.out

fmt:
	@echo "======================================================================"
	@echo "Run go fmt..."
	@go fmt ./...

clean_cache:
	@go clean -cache
	@go clean -testcache

mod-action-%:
	@echo "Run go mod ${*}...."
	GO111MODULE=on go mod $*
	@echo "Done go mod  ${*}"

mod: mod-action-verify mod-action-tidy mod-action-vendor mod-action-download mod-action-verify ## Download all dependencies
	@echo "All installed"
