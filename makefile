#!bash

migrate: #migration with goose: github.com/pressly/goose
	goose -dir ./migrations postgres "host=localhost sslmode=disable password=postgres dbname=packages user=postgres port=9948" up
update_docs: #documentation swagger with swaggo: github.com/swaggo/swag
	swag init -g ./cmd/server/main.go
lint: #golang-ci-lint: github.com/golangci/golangci-lint
	gofmt -s -w .
	golangci-lint run --timeout 5m0s -E gocyclo -E funlen -E lll ./...
tests:
	go test ./...
run_docker_db:
	docker compose -f ./config/docker/database.yaml up -d
run_docker_app:
	docker compose -f ./config/docker/app.yaml up -d
