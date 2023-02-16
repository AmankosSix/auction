#!make
include .env

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

postgres:
	docker-compose up -d postgresql

bash:
	docker exec -it pg-docker /bin/bash

run: build
	docker-compose up --remove-orphans app

migrate-up:
	migrate -path ${POSTGRES_MIGRATION_PATH} -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:5432/${POSTGRES_DBNAME}?sslmode=${POSTGRES_SSL}' up

migrate-down:
	migrate -path ${POSTGRES_MIGRATION_PATH} -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:5432/${POSTGRES_DBNAME}?sslmode=${POSTGRES_SSL}' down

migrate-init:
	migrate create -ext sql -dir ${POSTGRES_MIGRATION_PATH} -seq init

swag:
	swag init -g internal/app/app.go