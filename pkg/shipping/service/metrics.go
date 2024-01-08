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

	carrierList := make(map[string]shipping.CarrierMetrics)
	lowestPrice, higherPrice := math.MaxFloat64, -math.MaxFloat64

	for _, carrier := range carriers {
		if _, ok := carrierList[carrier.Name]; !ok {
			carrierList[carrier.Name] = shipping.CarrierMetrics{
				Total:        carrier.Price,
				AveragePrice: carrier.Price,
				Results:      1,
			}
			continue
		}

		carrierList[carrier.Name] = shipping.CarrierMetrics{
			Total:        carrierList[carrier.Name].Total + carrier.Price,
			AveragePrice: calculeAverage(carrierList[carrier.Name].Total, carrierList[carrier.Name].Results),
			Results:      carrierList[carrier.Name].Results + 1,
		}
	}

	for _, carrier := range carrierList {
		higherPrice = math.Max(higherPrice, carrier.Total)
		lowestPrice = math.Min(lowestPrice, carrier.Total)
	}

	metrics := &shipping.MetricsResponse{
		Carriers:    carrierList,
		HigherPrice: higherPrice,
		LowestPrice: lowestPrice,
	}

	return metrics, nil
}

func calculeAverage(total float64, resultsCount int64) float64 {
	return total / float64(resultsCount)
}
