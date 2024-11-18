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
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BUILD_DIR)/$(APP_NAME) -ldflags="$(LD_FLAGS)" $(SOURCE_DIR)

# Build for all platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		OUTPUT=$(BUILD_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH; \
		if [ "$$GOOS" = "windows" ]; then OUTPUT=$$OUTPUT.exe; fi; \
		echo "Building for $$platform..."; \
		CGO_ENABLED=$(CGO_ENABLED) GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$OUTPUT -ldflags="$(LD_FLAGS)" $(SOURCE_DIR); \
		echo "Built: $$OUTPUT"; \
	done

# Clean the build directory
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
	@echo "Cleaned."
