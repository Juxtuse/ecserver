include .env
.PHONY: help

help: ## show help msg
	@echo "Available commands:"
	@awk '/^[a-zA-Z0-9_-]+:([^=]|$$)/ { \
		split($$0, parts, ":"); \
		name = parts[1]; \
		if (match($$0, /## (.*)/, m)) { \
			printf "  \033[36m%-20s\033[0m %s\n", name, m[1]; \
		} else { \
			printf "  \033[36m%-20s\033[0m\n", name; \
		} \
	}' Makefile

dev: ## Run dev server
	reflex -d none -r '\.go$$' -s -- go run ./cmd/app/main.go

build: ## Build app
	go build -o build/app.out cmd/app/main.go

clean: ## Clean build folder
	rm -rf ./build

new-migration: ## Create new goose migration file
	cd migrations/$(or $(db), postgres) && goose create $(name) $(or $(type), sql) -s

fix-migration: ## Fix goose migration file
	cd migrations/$(or $(db), postgres) && goose fix

db-up: ## Run goose up
	cd migrations/postgres && goose postgres $(DATABASE_URL) -table go_backend_migrations up

db-up-by-one: ## Run goose up-by-one
	cd migrations/postgres && goose postgres $(DATABASE_URL) -table go_backend_migrations up-by-one

db-up-to: ## Run goose up-to [number]
	cd migrations/postgres && goose postgres $(DATABASE_URL) -table go_backend_migrations up-to $(version)

db-status: ## goose migration status
	cd migrations/postgres && goose postgres $(DATABASE_URL) -table go_backend_migrations status

db-down: ## Run goose down (it will only down one by one)
	cd migrations/postgres && goose postgres $(DATABASE_URL) -table go_backend_migrations down

db-down-to: ## Run goose down-to [number]
	cd migrations/postgres && goose postgres $(DATABASE_URL) -table go_backend_migrations down-to $(version)

init_backend:
	docker run --name valkey -d -p 6379:6379 valkey/valkey:8.1.1-alpine3.21
	docker run --name postgis-18 -p 5432:5432 -d -e POSTGRES_PASSWORD=mysecretpassword postgis/postgis:18-3.6-alpine

start-docker-services:
	docker start postgis-18 valkey

