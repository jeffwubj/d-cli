ifeq (, $(shell which go))
$(error Install Go - https://golang.org/dl/ )
endif

INSTALL_DIR=/usr/local/bin
BIN_DIR=./bin
NATIVE_ARCH=$(shell uname | tr A-Z a-z)

GOARCH=amd64
OSES=linux darwin windows
BUILD_TARGETS=$(foreach os,$(OSES),$(BIN_DIR)/$(os)/kubectl-docker)

.PHONY: clean
clean:
	-rm -rf $(BIN_DIR)

.PHONY: dist
dist: clean $(BUILD_TARGETS)

$(BIN_DIR)/%/kubectl-docker:
	GOOS=$* go build -o $@ ./cmd/kubectl-docker

.PHONY: build
build: $(BIN_DIR)/$(NATIVE_ARCH)/kubectl-docker 

$(BIN_DIR)/$(NATIVE_ARCH)/kubectl-docker:
	@go build -o $@ ./cmd/kubectl-docker

.PHONY: install
install: $(BIN_DIR)/$(NATIVE_ARCH)/kubectl-docker 
	cp $(BIN_DIR)/$(NATIVE_ARCH)/kubectl-docker $(INSTALL_DIR)

.PHONY: all
all: clean build install
