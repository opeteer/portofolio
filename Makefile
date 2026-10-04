.PHONY: all build install run clean test

BINARY_NAME=opeteer

all: build

build:
	@echo "Building opeteer CLI..."
	go build -o $(BINARY_NAME) .

install: build
	@echo "Installing $(BINARY_NAME) to $(shell go env GOPATH)/bin..."
	cp $(BINARY_NAME) $(shell go env GOPATH)/bin/$(BINARY_NAME)
	@echo "Installed! You can now run '$(BINARY_NAME)' from anywhere."

run: build
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

test: build
	@echo "Running Go unit tests..."
	go test -v ./...
	@echo "Verifying CLI commands..."
	./$(BINARY_NAME) --version
	./$(BINARY_NAME) about
	./$(BINARY_NAME) flagship
	./$(BINARY_NAME) showtime
	./$(BINARY_NAME) skills
	./$(BINARY_NAME) stats
	./$(BINARY_NAME) live --limit 2
	./$(BINARY_NAME) bench --target cpu --duration 300ms
	./$(BINARY_NAME) help live
	./$(BINARY_NAME) help bench
	./$(BINARY_NAME) help serve
