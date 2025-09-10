.PHONY: dev build run test tidy clean

# Start server in dev mode with Air (hot reload)
dev:
	air

# Build production binary
build:
	go build -o bin/go-saas-kit ./cmd/app

# Run production binary
run:
	./bin/go-saas-kit

# Run all tests
test:
	go test ./... -v

# Clean & tidy dependencies
tidy:
	go mod tidy
	go mod vendor

# Remove build artifacts
clean:
	rm -rf bin tmp
