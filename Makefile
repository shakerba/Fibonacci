# This is the Makefile for fibonnaci server

BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILD_ROOT=${PWD}
BUILDINFO_FILE=$(BUILD_ROOT)/build/fibonacci/buildinfo

.PHONY: test
export COVERAGE_DIR
test:
	./scripts/run-unit-tests

.PHONY: clean
clean:
	@echo "Cleaning up build/binaries"
	rm -rf $(BUILD_ROOT)/build && rm -rf $(BUILD_ROOT)/coverage

.PHONY: build
build: clean fibonacci
	@echo "Moving files to build directory"
	echo "{\"BUILD_TIME\":\"$(BUILD_TIME)\"}" > $(BUILD_ROOT)/build/docker-args.json
	echo > $(BUILDINFO_FILE)

.PHONY: fibonacci
fibonacci:
	@echo "Building fibonacci source code..."
	go version
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o $(BUILD_ROOT)/build/fibonacci/fibonacci github.com/pex/fibonacci/cmd/fibonacci
