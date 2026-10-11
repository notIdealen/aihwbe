include .env
export CURRENT_UID=$(shell id -u)
export CURRENT_GID=$(shell id -g)
export PROJECT_ROOT=$(shell pwd)

env-create-dirs:
	@mkdir -p "${PROJECT_ROOT}/${PSQL_DIR}"
	@mkdir -p "${PROJECT_ROOT}/${MIGRATE_DIR}"

# POSTGRES_SERVICE
POSTGRES_SERVICE_NAME:=psql
POSTGRES_CONTAINER_NAME:=psql-aihw
psql-up:
	docker compose up -d ${POSTGRES_SERVICE_NAME}
psql-down:
	docker compose down ${POSTGRES_SERVICE_NAME}
psql-delete:
	@make psql-down && \
	rm -rf ./out/pgdata/* && \
	echo "files inside pgdata deleted"
psql-connect:
	docker exec -it ${POSTGRES_CONTAINER_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

# MIGRATE_SERVICE
MIGRATE_SERVICE_NAME:=migrate
MIGRATE_CONTAINER_NAME:=migrate-aihw
MIGRATE_CMD = docker compose run --rm ${MIGRATE_SERVICE_NAME} \
	-path migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_CONTAINER_NAME}:5432/${POSTGRES_DB}?sslmode=disable"
migrate-create:
	@if [ -z "$(name)" ]; then \
	echo "Missing name, currect example: make migrate-create name=<name>"; \
	exit 1; \
	fi; \
	docker compose run --rm ${MIGRATE_SERVICE_NAME} \
	create \
	-ext sql \
	-dir migrations/ \
	-seq \
	$(name)
migrate:
ifdef N
	$(MIGRATE_CMD) $(act) $(N)
else
	$(MIGRATE_CMD) $(act)
endif
migrate-up:
	$(MAKE) migrate act=up
migrate-down:
	$(MAKE) migrate act=down
migrate-force-0:
	$(MAKE) migrate act=force N=0
migrate-version:
	$(MAKE) migrate act=version

# APP
app-run:
	go run ./cmd/main.go
