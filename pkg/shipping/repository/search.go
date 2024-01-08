package repository

import (
	"database/sql"

	"frete-rapido-api/pkg/shipping"

	"github.com/doug-martin/goqu/v9"
)

func (r *repository) Search(filter *shipping.Filter) ([]*shipping.CarrierDBModel, error) {
	statement, err := fullQueryCarriers(filter)
	if err != nil {
		return nil, err
	}

	rows, err := r.connection.Query(statement)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	carriers, err := iterateCarriers(rows)
	if err != nil {
		return nil, err
	}

	return carriers, nil
}

func iterateCarriers(rows *sql.Rows) ([]*shipping.CarrierDBModel, error) {
	carriers := make([]*shipping.CarrierDBModel, 0)

	for rows.Next() {
		var carrier shipping.CarrierDBModel
		err := rows.Scan(
			&carrier.ID,
			&carrier.Name,
			&carrier.Service,
			&carrier.Deadline,
			&carrier.Price,
		)
		if err != nil {
			return nil, err
		}

		carriers = append(carriers, &carrier)
	}

	return carriers, nil
}

func fullQueryCarriers(filter *shipping.Filter) (string, error) {
	dataset := goqu.From("carriers").Select("*")

	if filter.LastQuotes != 0 {
		dataset = dataset.Limit(uint(filter.LastQuotes))
	}

	statement, _, err := dataset.ToSQL()
	if err != nil {
		return "", err
	}

	return statement, nil
}
