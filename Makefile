BINARY ?= go-cryptotool.exe

.PHONY: run build clean test help

run:
	go run ./cmd/go-cryptotool/main.go

build:
	go build -o $(BINARY) ./cmd/go-cryptotool/main.go 

clean:
	rm -f ./$(BINARY)

test:
	go test ./tests/... -v

help:
	@echo "Commands:"
	@echo "  make run   - Runs the CLI for testing purposes"
	@echo "  make build - Builds the CLI into a binary"
	@echo "  make clean - Removes the built CLI binary"
	@echo "  make test  - Runs the Ciphers API tests"