package service

import (
	"fmt"
	"math"
	"strconv"

	"shipping-calculator-api/pkg/shipping"
)

func (s *service) Metrics(filter *shipping.Filter) (*shipping.MetricsResponse, error) {
	carriers, err := s.shippingRepo.Search(filter)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrQuery, err)
	}

	if len(carriers) < 1 {
		return nil, ErrNotFound
	}

	carrierMap := make(map[string]shipping.CarrierMetrics)
	lowestPrice, higherPrice := math.MaxFloat64, -math.MaxFloat64

	if err := calcMetrics(carrierMap, carriers); err != nil {
		return nil, err
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

func calcMetrics(
	carrierMap map[string]shipping.CarrierMetrics,
	carriers []*shipping.CarrierDBModel,
) error {
	for _, carrier := range carriers {
		if _, ok := carrierMap[carrier.Name]; !ok {
			carrierMap[carrier.Name] = shipping.CarrierMetrics{
				Total:        carrier.Price,
				AveragePrice: carrier.Price,
				Results:      1,
			}
			continue
		}

		totalFormated := fmt.Sprintf(
			"%.2f",
			carrierMap[carrier.Name].Total+carrier.Price,
		)
		totalConverted, err := strconv.ParseFloat(totalFormated, 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrConvertString, err)
		}

		averageFormated := fmt.Sprintf(
			"%.2f",
			calculeAverage(carrierMap[carrier.Name].Total, carrierMap[carrier.Name].Results),
		)
		averageConverted, err := strconv.ParseFloat(averageFormated, 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrConvertString, err)
		}

		results := carrierMap[carrier.Name].Results + 1
		carrierMap[carrier.Name] = shipping.CarrierMetrics{
			Total:        totalConverted,
			AveragePrice: averageConverted,
			Results:      results,
		}
	}

	return nil
}

func calculeAverage(total float64, resultsCount int64) float64 {
	return total / float64(resultsCount)
}
