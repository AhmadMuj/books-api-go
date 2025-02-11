.PHONY: build run test clean swagger dev docker-dev docker-prod docker-down

BINARY_NAME=books-api-go
BUILD_DIR=build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/api/main.go

run:
	./$(BUILD_DIR)/$(BINARY_NAME)

dev:
	air -c .air.toml --build.poll=true

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)

swagger:
	swag init -g ./cmd/api/main.go -o ./docs/swagger

docker-dev:
	docker-compose -f deployments/docker-compose.dev.yml up --build

docker-prod:
	docker-compose -f deployments/docker-compose.prod.yml up --build -d

docker-down:
	docker-compose -f deployments/docker-compose.dev.yml down
	docker-compose -f deployments/docker-compose.prod.yml down
	
.DEFAULT_GOAL := build
