SHELL := /usr/bin/env bash

.EXPORT_ALL_VARIABLES:
OUT_DIR ?= _output
BIN_DIR := $(OUT_DIR)/bin
MODULES_DIR := $(BIN_DIR)/modules
DEBUG_DIR :=$(OUT_DIR)/debug
PLUGIN_DIR :=$(OUT_DIR)/plugins
RELEASE_DIR :=$(OUT_DIR)/release
PACKAGE_DIR :=build/package
RELEASE_TAR_DIR := $(OUT_DIR)/release_tar

define ALL_HELP_INFO
# Build code.
#
# Example:
# make run 
# make test  --单元测试
endef

.PHONY: all
all: clean test run

.PHONY: run
run: 
	go run cmd/holdem/main.go

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:

.PHONY: help
help:
	@echo "$$ALL_HELP_INFO"



