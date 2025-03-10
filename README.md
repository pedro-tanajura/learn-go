# Learn Go Repository

This repository contains various Go code examples and study materials organized in different directories.

## Repository Structure

- `data-structures/`: Implementation of various data structures
  - `stack/`: Stack implementation
  - `queue/`: Queue implementation
  - `singlyLinkedList/`: Singly linked list implementation
  - `doublyLinkedList/`: Doubly linked list implementation
  - `minStack/`: Min stack implementation
- `greetings/`: A simple greeting module
- `hello/`: A hello world application that uses the greetings module

## Running Code

You can run any program in this repository using one of the following methods:

### Using the Makefile (Recommended)

```bash
# List all available programs
make list

# Run a specific program
make run PROGRAM=data-structures/stack
```

### Using the Shell Script

```bash
# Run a specific program
./run.sh data-structures/stack
```

### Using the Go Runner

```bash
# Run a specific program
go run run.go data-structures/stack
```

## Adding New Code

To add new code to this repository:

1. Create a new directory for your code
2. Add your Go files to the directory
3. Optionally, create a `go.mod` file if your code is a module

Your new code will automatically be available to run using the methods above.

## Examples

```bash
# Run the stack implementation
make run PROGRAM=data-structures/stack

# Run the hello world application
make run PROGRAM=hello
```
