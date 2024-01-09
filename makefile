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