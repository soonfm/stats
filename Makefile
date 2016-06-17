#
# SFM 2.0 Stats Service
#

VERSION 	?= $(shell git rev-parse --short HEAD)
BUILD_DIR 	?= _build
BUILD_NAME 	?= playlist
BUILD_FLAGS ?= -ldflags "-X main.VERSION=${VERSION}"

install_dependencies:
	govendor sync

#
# Build all binaries
#

build: install_dependencies build_linux build_osx

#
# Linux Build
#

build_linux: build_linux64

build_linux64:
	mkdir -p ${BUILD_DIR}/linux_amd64
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/linux_amd64/$(BUILD_NAME) ${BUILD_FLAGS}

#
# OSX Build
#

build_osx: build_osx64

build_osx64:
	mkdir -p ${BUILD_DIR}/darwin_amd64
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/darwin_amd64/$(BUILD_NAME) ${BUILD_FLAGS}

#
# Testing
#

test:
	go test -v -cover ./...
