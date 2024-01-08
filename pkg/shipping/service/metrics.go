package service

import (
	"math"

	"frete-rapido-api/pkg/shipping"
)

func (s *service) Metrics(filter *shipping.Filter) (*shipping.MetricsResponse, error) {
	carriers, err := s.shippingRepo.Search(filter)
	if err != nil {
		return nil, err
	}

	if len(carriers) < 1 {
		return nil, ErrNotFound
	}

	carrierMap := make(map[string]shipping.CarrierMetrics)
	lowestPrice, higherPrice := math.MaxFloat64, -math.MaxFloat64

	for _, carrier := range carriers {
		if _, ok := carrierMap[carrier.Name]; !ok {
			carrierMap[carrier.Name] = shipping.CarrierMetrics{
				Total:        carrier.Price,
				AveragePrice: carrier.Price,
				Results:      1,
			}
			continue
		}

		total := carrierMap[carrier.Name].Total + carrier.Price
		average := calculeAverage(carrierMap[carrier.Name].Total, carrierMap[carrier.Name].Results)
		results := carrierMap[carrier.Name].Results + 1

		carrierMap[carrier.Name] = shipping.CarrierMetrics{
			Total:        total,
			AveragePrice: average,
			Results:      results,
		}
	}

	for _, carrier := range carrierMap {
		higherPrice = math.Max(higherPrice, carrier.Total)
		lowestPrice = math.Min(lowestPrice, carrier.Total)
	}

	metrics := &shipping.MetricsResponse{
		Carriers:    carrierMap,
		LowestPrice: lowestPrice,
		HigherPrice: higherPrice,
	}

	return metrics, nil
}

func calculeAverage(total float64, resultsCount int64) float64 {
	return total / float64(resultsCount)
}
