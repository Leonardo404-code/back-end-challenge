//go:build wireinject
// +build wireinject

package handler

import (
	"database/sql"
	"shipping-calculator-api/pkg/shipping"
	"shipping-calculator-api/pkg/shipping/repository"
	"shipping-calculator-api/pkg/shipping/service"

	"github.com/google/wire"
)

func Build(db *sql.DB) shipping.Handler {
	wire.Build(
		Must,
		service.Must,
		repository.Must,
	)

	return nil
}
