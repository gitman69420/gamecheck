# Made with ❤️, but with AI 🤖

.PHONY: help dev dev-down dev-logs prod prod-down prod-logs sqlc migrate-dev migrate-prod migrate-create migrate-status clean db-seed

# Load environment variables
include .env.local
export

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development Commands
dev: ## Start development environment with hot reload
	docker-compose -f docker-compose.dev.yaml up --build

dev-down: ## Stop development environment
@RUNNING=$$(docker-compose -f docker-compose.dev.yaml ps --status running -q 2>/dev/null); \
if [ -n "$$RUNNING" ]; then \
	echo "Stopping containers..."; \
	docker-compose -f docker-compose.dev.yaml down; \
else \
	echo "No running containers, skipping down command"; \
fi

dev-logs: ## View development logs
	docker-compose -f docker-compose.dev.yaml logs -f

dev-rebuild: ## Rebuild and restart development environment
	docker-compose -f docker-compose.dev.yaml down
	docker-compose -f docker-compose.dev.yaml up --build

# # Production Commands
# prod: ## Build and start production environment
# 	docker-compose up --build -d

# prod-down: ## Stop production environment
# 	docker-compose down

# prod-logs: ## View production logs
# 	docker-compose logs -f

# prod-rebuild: ## Rebuild production environment
# 	docker-compose down
# 	docker-compose up --build -d

# Database Migration Commands
migrate-dev: ## Run database migrations (development)
	docker-compose -f docker-compose.dev.yaml run --rm atlas \
		migrate apply \
		--dir "file:///migrations" \
		--url "postgres://$(DB_USER):$(DB_PASSWORD)@postgres:5432/$(DB_NAME)?sslmode=disable"

# migrate-prod: ## Run database migrations (production)
# 	docker-compose run --rm atlas \
# 		migrate apply \
# 		--dir "file:///migrations" \
# 		--url "postgres://$(DB_USER):$(DB_PASSWORD)@postgres:5432/$(DB_NAME)?sslmode=disable"

migrate-create: ## Create a new migration from postgres-schema.sql (usage: make migrate-create NAME=migration_name)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	docker-compose -f docker-compose.dev.yaml run --rm atlas \
		migrate diff $(NAME) \
		--dir "file:///migrations" \
		--to "file:///postgres-schema.sql" \
		--dev-url "postgres://postgres:mysecretpassword@dev-atlas-devdb:5432/postgres?sslmode=disable"

migrate-status: ## Check migration status (development)
	docker-compose -f docker-compose.dev.yaml run --rm atlas \
		migrate status \
		--dir "file:///migrations" \
		--url "postgres://$(DB_USER):$(DB_PASSWORD)@postgres:5432/$(DB_NAME)?sslmode=disable"

migrate-hash: ## Rehash migration directory
	docker-compose -f docker-compose.dev.yaml run --rm atlas \
		migrate hash \
		--dir "file:///migrations"

# SQLC Commands
sqlc: ## Generate Go code from SQL queries
	docker-compose -f docker-compose.dev.yaml exec app sqlc generate

sqlc-local: ## Generate Go code from SQL queries (local - without docker)
	sqlc generate

# Utility Commands
clean: ## Clean up Docker resources
	docker-compose -f docker-compose.dev.yaml down -v
	docker-compose down -v
	docker system prune -f

db-shell: ## Connect to PostgreSQL database (development)
	docker-compose -f docker-compose.dev.yaml exec postgres psql -U $(DB_USER) -d $(DB_NAME)

valkey-cli: ## Connect to Valkey CLI (development)
	docker-compose -f docker-compose.dev.yaml exec valkey valkey-cli

test: ## Run tests in Docker
	docker-compose -f docker-compose.dev.yaml exec app go test ./...

test-local: ## Run tests locally
	go test ./...