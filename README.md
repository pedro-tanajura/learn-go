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

## Using the Alias (Easiest)

A convenient `gorun` script is included in this repository. You can set it up as an alias or add it to your PATH for easier access.

```bash
# After setting up the alias as described in alias_setup.txt
gorun data-structures/stack
gorun hello
```

See the `alias_setup.txt` file for instructions on how to set up the alias.

## Examples

```bash
# Run the stack implementation
make run PROGRAM=data-structures/stack
gorun data-structures/stack  # Using the alias after setup

# Run the hello world application
make run PROGRAM=hello
gorun hello  # Using the alias after setup
```
