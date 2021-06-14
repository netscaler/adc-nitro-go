SDK_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/" | grep -v "/samples")
SDK_TEST_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/" | grep -v "/config" | grep -v "/stat"| grep -v "/samples")

all: build unit

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build                   to go build the SDK"
	@echo "  unit                    to run unit tests"
	@echo "  lint                    to lint the SDK"

build:
	@echo "go build SDK and vendor packages"
	@go build ${SDK_ONLY_PKGS}

unit:  build 
	@echo "go test SDK  package"
	@go test  -v $(SDK_TEST_ONLY_PKGS)

lint:  build 
	@echo "go lint service package (ignoring generated packages)"
	@golint   service | grep -v netscaler/resources.go || true

