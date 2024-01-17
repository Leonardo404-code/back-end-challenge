package repository

import (
	"database/sql"
	"sync"

	"shipping-calculator-api/pkg/shipping"
)

type repository struct {
	connection *sql.DB
}

var (
	repo shipping.Repository
	once sync.Once
)

func New(opts ...Option) (shipping.Repository, error) {
	r := &repository{}

	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func Must(connection *sql.DB) shipping.Repository {
	once.Do(func() {
		r, err := New(WithConnection(connection))
		if err != nil {
			panic(err)
		}

		repo = r
	})

	return repo
}
