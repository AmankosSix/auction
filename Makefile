.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

runDB:
	docker-compose up -d postgresql

bash:
	docker exec -it pg-docker /bin/bash

run: build
	docker-compose up --remove-orphans app