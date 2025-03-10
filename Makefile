# Makefile for learn-go repository

.PHONY: run help list

# Default target
help:
	@echo "Usage:"
	@echo "  make run PROGRAM=<program-path>  - Run a specific program"
	@echo "  make list                       - List all available programs"
	@echo "  make help                       - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make run PROGRAM=data-structures/stack"
	@echo "  make run PROGRAM=hello"

# Run a program
run:
	@if [ -z "$(PROGRAM)" ]; then \
		echo "Error: PROGRAM parameter is required"; \
		echo "Usage: make run PROGRAM=<program-path>"; \
		exit 1; \
	fi
	@./run.sh $(PROGRAM)

# List all available programs
list:
	@echo "Available programs:"
	@find . -mindepth 1 -maxdepth 3 -type d -not -path "*/\.*" | sort | sed 's|^\./||' | \
		while read dir; do \
			if [ -f "$$dir/go.mod" ] || [ -n "$$(find "$$dir" -maxdepth 1 -name "*.go" 2>/dev/null)" ]; then \
				echo "  $$dir"; \
			fi; \
		done
