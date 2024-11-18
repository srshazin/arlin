# Project settings
APP_NAME := arlin
BUILD_DIR := build
SOURCE_DIR := .
GO_FILES := $(shell find $(SOURCE_DIR) -type f -name '*.go')

# Default targets for architectures
PLATFORMS := linux/amd64 linux/arm64

# Optimization flags
LD_FLAGS := -s -w
CGO_ENABLED := 0

.PHONY: all clean build build-all

# Default target
all: build

# Build for the current platform
build:
	@echo "Building for the current platform..."
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)  .

# Clean the build directory
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
	@echo "Cleaned."
