include .env
export CURRENT_UID=$(shell id -u)
export CURRENT_GID=$(shell id -g)
export PROJECT_ROOT=$(shell pwd)

env-create-dirs:
	@mkdir -p "${PROJECT_ROOT}/out/pgdata"
	@mkdir -p "${PROJECT_ROOT}/migrations"

psql-up:
	docker compose up -d hw3-psql
psql-down:
	docker compose down hw3-psql
psql-delete:
	@make psql-down && \
	rm -rf ./out/pgdata/* && \
	echo "files inside pgdata deleted"

migrate-create:
	@if [ -z "$(name)" ]; then \
	echo "Missing name, currect example: make migrate-create name=<name>"; \
	exit 1; \
	fi; \
	docker compose run --rm hw3-migrate \
	create \
	-ext sql \
	-dir migrations/ \
	-seq \
	$(name)

migrate-up:
	docker compose run --rm hw3-migrate \
	-path migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@hw3-psql:5432/${POSTGRES_DB}?sslmode=disable" \
	up

migrate-down:
	docker compose run --rm hw3-migrate \
	-path migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@hw3-psql:5432/${POSTGRES_DB}?sslmode=disable" \
	down

app-run:
	go run ./cmd/main.go
