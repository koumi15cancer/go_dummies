run: build
	@./bin/api

build:
	@go build -o bin/api

test:
	@go test -v ./...

lean:
	@rm -rf bin
	@echo "Cleaned build artifacts."

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down