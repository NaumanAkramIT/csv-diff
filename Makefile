.PHONY: build test test-race run clean help

build:
	go build -o csv-diff ./app/cmd

test:
	go test ./...

test-race:
	go test -race -cover ./...

run: build
	./csv-diff -c ./sample/file1.csv ./sample/file2.csv

clean:
	rm -f csv-diff

help:
	@echo "Available targets:"
	@echo "  build       Build the csv-diff binary"
	@echo "  test        Run unit tests"
	@echo "  test-race   Run tests with race detector and coverage"
	@echo "  run         Build and run csv-diff with sample files"
	@echo "  clean       Remove built binaries"