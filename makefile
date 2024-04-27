# Makefile 示例

# Compiler (例如 gcc，g++，或 go build，视你的项目而定)
CC := go build

# Source files
SRC := $(wildcard *.go)

# Output binary name
BINARY := myapp

# Build the project
# This target builds the project by compiling the source files and generating the binary
build: $(SRC)
    $(CC) -o $(BINARY) $(SRC)

# Clean built binary
# This target cleans the built binary by removing it
clean:
    rm -f $(BINARY)

# Run tests
# This target runs the tests in the project
test:
    go test ./...