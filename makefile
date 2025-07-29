# Packages to test (excluding those under /test directories)
PKGS := $(shell go list ./... | grep -vE '/(test|docs)')

# Run unit tests with verbose output
test:
	go test -v $(PKGS)

# Generate test coverage profile
coverage:
	go test -cover -coverprofile=coverage.out $(PKGS)

# Generate HTML report from coverage
coverage-html: coverage
	go tool cover -html=coverage.out

# Show total coverage in terminal
coverage-total: coverage
	go tool cover -func=coverage.out | grep total | awk '{print "total coverage: " $$3}'

# Run static analysis (linter)
lint:
	staticcheck ./...

# Clean up generated files
clean:
	rm -f coverage.out coverage.html

.PHONY: test coverage coverage-html coverage-total lint clean
