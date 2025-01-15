build:
	@echo "Building application as a binary 'JammeApp'..."
	@go build -o ./bin/JammeApp .
test:
	@echo "Running tests..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html coverage.out -o coverage.html

lint:
	@golangci-lint run

docs:
	@swag fmt && swag init --parseDependency --parseInternal

dev:
	@cls && set ENVIRONMENT=development && go run main.go

docker-build:
	@docker build -t  azarm98/pet-store:latest .

docker-ghcr-build:
	@docker build -t azarm98/pet-store:latest .

fmt:
	@go fmt ./...

docker-push:
	@docker push azarm98/pet-store:latest 
	
.PHONY: test docs dev docker-build docker-push fmt lint build