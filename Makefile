include .env
export

GO_RELOAD = $(GOPATH)/bin/reflex -s -r '\.go$$' --

build:
	@go build -tags $(GO_TAGS) -o ./bin/ms-transaction ./cmd/transaction

infra.up:
	@docker-compose up -d --remove-orphans postgres

run-api:
	@$(GO_RELOAD) go run -tags $(run_cmd) -race ./cmd/main.go api