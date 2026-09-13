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

app-run:
	go run ./cmd/main.go
