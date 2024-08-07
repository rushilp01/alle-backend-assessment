.PHONY: all run clean

# The default target that runs when you run `make` with no arguments
all: gorun

# Run the Go application
gorun:
	go run .

# Clean up any build artifacts (none in this simple example, but useful for future expansion)
clean:
	go clean

# Ensure dependencies are installed
deps:
	go mod tidy

# Shortcut to install dependencies and run the application
gorun: deps run
