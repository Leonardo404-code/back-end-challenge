//go:build wireinject
// +build wireinject

package handler

import (
	"database/sql"
	"frete-rapido-api/pkg/shipping"

	"github.com/google/wire"
)

func Build(db *sql.DB) shipping.Handler {
	wire.Build(
		Must,
	)

	return nil
}
