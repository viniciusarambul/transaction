include .env
export

build:
	@go build -tags $(run_cmd) -o ./bin/ms-transaction src/api/

infra.up:
	@docker-compose up -d --remove-orphans postgres

run-api:
	@go run -tags $(run_cmd) -race src/api/main.go

flyway:
	@docker-compose up flyway