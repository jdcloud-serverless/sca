# Build binary
#
# Example:
#   make
#   make all
all: build

# Build binary
#
# Example:
#   make build
build:
	go build -v -o ./bin/sca .

# Generate binary vendor
#
# Example:
#   make generate
generate:
	go mod vendor

# Clean package
#
# Example:
#   make clean
clean:
	rm -rf ./bin

