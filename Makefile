.PHONY: help postgres createdb dropdb migratedown migrateup sqlc test server mock install-tools

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}'

install-tools: ## Install dev-tools for project
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install go.uber.org/mock/mockgen@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

createNetwork: ## Create Docker network for the application
	docker network create bank-network

run: ## Start the application using docker-compose
	docker-compose up -d

postgres: ## Start PostgreSQL container
	docker run -d \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=simple_bank \
		-p 5432:5432 \
		--name postgres12 \
		postgres:17.2

createdb: ## Create a new database
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb: ## Drop the database
	docker exec -it postgres12 dropdb simple_bank

migrateup: ## Run all up database migrations
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1: ## Run one up database migration
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown: ## Run all down database migrations
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1: ## Run one down database migration
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc: ## Generate Go code from SQL
	sqlc generate

test: ## Run all unit tests
	go test -v -cover ./...

server: ## Start the application server
	go run main.go

dockerbuild: ## Build Docker image
	docker build -t simplebank:latest .

dockerrun: ## Run application in Docker container
	docker run --name simplebank \
		-p 8080:8080 \
		--network bank-network \
		-e GIN_MODE=release \
		-e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" \
		simplebank:latest

mock: ## Generate mock store for testing
	mockgen -package mockdb -destination db/mock/store.go github.com/cs-tungthanh/Bank_Golang/db/sqlc Store
