package repository

import (
	"fmt"

	"shipping-calculator-api/pkg/shipping"

	"github.com/google/uuid"
)

func (r *repository) Create(carrier *shipping.CarrierInfo) error {
	shipppingDB := &shipping.CarrierDBModel{
		ID:       uuid.New().String(),
		Name:     carrier.Name,
		Service:  carrier.Service,
		Deadline: carrier.Deadline,
		Price:    carrier.Price,
	}

	insertResult, err := r.connection.Exec(
		`INSERT INTO carriers (id, name, service, deadline, price) VALUES ($1, $2, $3, $4, $5)`,
		shipppingDB.ID,
		shipppingDB.Name,
		shipppingDB.Service,
		shipppingDB.Deadline,
		shipppingDB.Price,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := insertResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in get rows affected")
	}

	if rowsAffected != 1 {
		return fmt.Errorf("no rows affected in database")
	}

	return nil
}
