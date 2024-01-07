package repository

import (
	"database/sql"
)

type Option func(*repository) error

func WithConnection(conn *sql.DB) Option {
	return func(r *repository) error {
		r.connection = conn
		return nil
	}
}
