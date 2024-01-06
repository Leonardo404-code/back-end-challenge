//go:build wireinject
// +build wireinject

package handler

import (
	"database/sql"

	"github.com/google/wire"

	"frete-rapido-api/pkg/shipping"
	"frete-rapido-api/pkg/shipping/repository"
	"frete-rapido-api/pkg/shipping/service"
)

func Build(db *sql.DB) shipping.Handler {
	wire.Build(
		Must,
		service.Must,
		repository.Must,
	)

	return nil
}
