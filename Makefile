# Makefile for Distribution System Program

# Define variables
PROGRAM_NAME := qube
GO_SRC_FILES := main.go

# Build the Go program
build:
	go build -o $(PROGRAM_NAME) $(GO_SRC_FILES)

# Run the Go program with the specified options
run:
	./$(PROGRAM_NAME) $(ARGS)

# Clean up the build artifacts
clean:
	rm -f $(PROGRAM_NAME)

# Run the program with the include and exclude flags
example:
	./$(PROGRAM_NAME) --include IN --exclude CN

.PHONY: build run clean example
