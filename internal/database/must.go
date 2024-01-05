package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"

	"frete-rapido-api/internal/env"
)

var (
	connection *sql.DB
	once       sync.Once
)

func Must() *sql.DB {
	once.Do(func() {
		dbURL := env.GetString("DATABASE.URL")

		postgresConn, err := sql.Open("pgx", dbURL)
		if err != nil {
			panic(err)
		}

		if err := postgresConn.Ping(); err != nil {
			panic(fmt.Errorf("failed to ping database: %w", err))
		}

		connection = postgresConn
	})

	return connection
}
