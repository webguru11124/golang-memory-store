.PHONY: all build run test docker-build docker-run clean

all: build

build:
	@echo "Building the application..."
	go build -o server ./cmd/server

run: build
	@echo "Running the application..."
	./server

test:
	@echo "Running unit tests..."
	go test -v ./internal/core
	@echo "Running integration tests..."
	go test -v ./tests

docker-build:
	@echo "Building Docker image..."
	docker build -t golang-memory-store .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 golang-memory-store

clean:
	@echo "Cleaning up..."
	rm -f server
