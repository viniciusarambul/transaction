include .env
export

build:
	@go build

infra.up:
	@docker-compose up -d --remove-orphans postgres

run-api:
	@go run -tags $(run_cmd) -race main.go

flyway:
	@docker-compose up flyway

test.cover:
	@go test ./... -coverprofile=cover.txt
	@go tool cover -html=cover.txt -o cover.html
	@open cover.html
	