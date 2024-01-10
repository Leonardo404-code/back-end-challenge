# Frete Rápido Back-end challenge

Rest API for external queries of delivery data in the Frete Rápido API

## API Documentation

See [swagger.yaml](/docs/swagger.yaml)

## Requirements

- Golang 18+
- Docker

## How to run

### Setup database

- Install Docker

Docker is a software platform that allows developers to create, deploy, and run applications in containers. See the official documentation for install the version compatible with your OS: [Install Docker Engine](https://docs.docker.com/engine/install/)

### Run Docker compose

```docker
docker compose -f ./config/docker/docker-compose.yaml -d
```

### Download the project dependencies

```go
go mod vendor
```

### Set environment variables
- Set environment variables listed in the table below

## Migrations
- Perform database migrations (with [goose](https://github.com/pressly/goose) or directly in the database)

## Execute the project

- Inside frete-rapido-api directory, run:

```go
go run cmd/server/main.go
```

- That's it, the application will be running locally on the port you configured! 🚀

## Running with docker

- Install docker

### Run Docker compose

```docker
docker compose -f ./config/docker/docker-compose.yaml -d
```

- List running containers

```docker
docker ps
```

You should see the application and database containers running, like:

![docker list containers output](/docs/docker-containers-output.png)

### Migrations

- Perform database migrations (with [goose](https://github.com/pressly/goose) or directly in the database)

- Application will be running locally on port 3000! 🚀

## Testes

- Inside frete-rapido-api directory, run:

```go
go test ./... -v
```

You should see messages like:

![test output](/docs/test-output.png)

## Environment Variables

| Name | Description | Example |
| --- | --- | --- |
| DATABASE_URL | Database connection URL | postgres://postgres:postgres@localhost:5432/packages |
| PORT | Port where the application listen for HTTP requests | 3000 |