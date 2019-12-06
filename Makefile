CURDIR := $(shell pwd)
GOBIN := $(CURDIR)/bin/
ENV:=GOBIN=$(GOBIN)
DIR:=FILE_DIR=$(CURDIR)/testfiles TEST_SOURCE_PATH=$(CURDIR)
GODEBUG:=GODEBUG=gocacheverify=1

##
## List of commands:
##

## default:
all: mod deps fmt lint test

all-deps: mod deps

deps:
	@echo "======================================================================"
	@echo 'MAKE: deps...'
	@mkdir -p $(GOBIN)
	@$(ENV) go get -u golang.org/x/lint/golint

test: tests

tests:
	@echo "Run test ./..."
	@$(DIR) $(GODEBUG) go test ./...

tests-cover:
	@echo "Run test  -cover -race -coverprofile=./coverage.out ./..."
	@$(DIR) $(GODEBUG) go test  -cover -race -coverprofile=./coverage.out ./...
	go tool cover -html=./coverage.out -o ./coverage.html
	rm ./coverage.out

lint:
	@echo "======================================================================"
	@echo "Run golint..."
	$(GOBIN)golint ./...

fmt:
	@echo "======================================================================"
	@echo "Run go fmt..."
	@go fmt ./...

mod:
	@echo "======================================================================"
	@echo "Run MOD"
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod verify
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod tidy
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod vendor
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod download
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod verify

clean_cache:
	@go clean -cache
	@go clean -testcache
