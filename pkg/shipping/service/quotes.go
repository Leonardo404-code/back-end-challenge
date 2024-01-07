package service

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"

	"frete-rapido-api/pkg/shipping"
)

// TODO: create repository for create role
// TODO: return data from handler
func (s *service) Quotes(
	shippingData *shipping.ShippingDataRequest,
) (*shipping.ShippingDataResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(shippingData); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	quotes, err := s.Fetch(shippingData)
	if err != nil {
		return nil, err
	}

	quotesData := new(shipping.QuotesData)

	if err := json.Unmarshal(quotes, quotesData); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	response := new(shipping.ShippingDataResponse)

	for _, data := range quotesData.Dispatcher[0].Offers {
		response.Carrier = append(response.Carrier, shipping.CarrierInfo{
			Name:     data.Carrier.Name,
			Service:  data.Service,
			Deadline: data.DeliveryTime.Days,
			Price:    data.FinalPrice,
		})
	}

	return response, nil
}
