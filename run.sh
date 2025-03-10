#!/bin/bash

# Check if a program path is provided
if [ $# -lt 1 ]; then
    echo "Usage: ./run.sh <program-path>"
    echo "Example: ./run.sh data-structures/stack"
    echo "Example: ./run.sh hello"
    exit 1
fi

# Get the program path from command line arguments
PROGRAM_PATH="$1"

# Get the absolute path to the repository root
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Construct the full path to the program
FULL_PATH="$REPO_ROOT/$PROGRAM_PATH"

# Check if the directory exists
if [ ! -d "$FULL_PATH" ]; then
    echo "Error: $PROGRAM_PATH is not a valid directory"
    exit 1
fi

# Check if there's a go.mod file in the directory
if [ -f "$FULL_PATH/go.mod" ]; then
    echo "Running program in $PROGRAM_PATH (with go.mod)"
    cd "$FULL_PATH" && go run .
    exit $?
fi

# Find all .go files in the directory
GO_FILES=$(find "$FULL_PATH" -maxdepth 1 -name "*.go")
if [ -z "$GO_FILES" ]; then
    echo "Error: No Go files found in $PROGRAM_PATH"
    exit 1
fi

echo "Running program in $PROGRAM_PATH"
cd "$FULL_PATH" && go run *.go
