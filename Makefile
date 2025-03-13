.PHONY: run build deps

# Путь к .env файлу
ENV_FILE := .env

# Цель для загрузки переменных из .env
load-env:
	@if [ -f $(ENV_FILE) ]; then \
		export $$(grep -v '^#' $(ENV_FILE) | xargs); \
	fi


# Формирование DB_MIGRATE_URL из переменных окружения
DB_MIGRATE_URL := postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable
MIGRATE_PATH := ./migrations

run:
	go run cmd/main.go

build:
	go build -o bin/testEx cmd/main.go

deps:
	go mod download

	migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create: load-env
	@if [ -z "$(name)" ]; then \
		echo "❌ Ошибка: укажите имя миграции с name=<название>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir "$(MIGRATE_PATH)" -seq "$(name)"



migrate-up: load-env
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down: load-env
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all
