#!/bin/bash

# Get the build path from command line arguments
if [ "$#" -eq 0 ]; then
  echo "Usage: build.sh <build-path>"
  exit 1
fi

BUILD_PATH=$1

# Clean the project
go clean "${BUILD_PATH}"

# Build the project
go build "${BUILD_PATH}/..."

# Tidy the Go modules
go mod tidy

