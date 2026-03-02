SHELL := /bin/sh

-include .env
export

APP_NAME := api
APP_PORT ?= 8080

DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= tefabene

DB_DSN := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

MIGRATIONS_DIR := db/migrations
SEEDS_DIR := db/seeds

# ---------- Help ----------
.PHONY: help
help:
	@echo "Targets:"
	@echo "  make up             - Runs both app and database (docker compose)"
	@echo "  make down           - Drops containers"
# 	@echo "  make logs           - App logs"
	@echo "  make db-logs        - Database logs"
	@echo "  make migrate-up     - Apply migrations (goose)"
	@echo "  make migrate-down   - Reverts 1 migration"
	@echo "  make migrate-status - Migration status"
	@echo "  make seed           - Runs seeds"
	@echo "  make run            - Runs API locally (go run)"
	@echo "  make build          - Builds locally"
	@echo "  make test           - Tests"
	@echo "  make fmt            - Go format"
	@echo "  make swaggo"        - Generates Swagger docs (swaggo)

# ---------- Docker ----------
.PHONY: up down restart logs db-logs ps swaggo
up:
	docker compose up -d --build

down:
	docker compose down

restart:
	docker compose down
	docker compose up -d --build

logs:
	docker compose logs -f app

db-logs:
	docker compose logs -f database

ps:
	docker compose ps

swaggo:
	swag init -g cmd/api/main.go -d . -o docs

# ---------- Local run/build ----------
.PHONY: run build test fmt
run:
	go run ./cmd/api/main.go

build:
	go build -o bin/$(APP_NAME) ./cmd/api/main.go

test:
	go test ./... -count=1

fmt:
	gofmt -w .

# ---------- Migrations (Goose) ----------
# Requires goose install: go install github.com/pressly/goose/v3/cmd/goose@latest
.PHONY: migrate-up migrate-down migrate-status migrate-new
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down 1

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" status

# use: make migrate-new name=create_users
migrate-new:
	@if [ -z "$(name)" ]; then echo "use: make migrate-new name=create_users"; exit 1; fi
	goose -dir $(MIGRATIONS_DIR) create "$(name)" sql

# ---------- Seeds ----------
.PHONY: seed
seed:
	@echo "Seeding database using ./cmd/seed/main.go..."
	go run ./cmd/seed/main.go
