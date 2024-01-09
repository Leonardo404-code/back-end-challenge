package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	connection *sql.DB
	once       sync.Once
)

func Must() *sql.DB {
	once.Do(func() {
		dbURL := os.Getenv("DATABASE_URL")

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
