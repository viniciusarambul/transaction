include .env
export

build:
	@go build -tags $(GO_TAGS) -o ./bin/ms-transaction ./cmd/transaction

infra.up:
	@docker-compose up -d --remove-orphans postgres app

run-api:
	@go run -tags $(run_cmd) -race src/api/main.go