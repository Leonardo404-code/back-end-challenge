package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"shipping-calculator-api/pkg/shipping"
)

func (s *service) Fetch(shippingData *shipping.ShippingDataRequest) ([]byte, error) {
	convertToJSON, err := json.Marshal(shippingData)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMarshal, err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		apiURL,
		bytes.NewBuffer(convertToJSON),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateRequest, err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrPerformRequest, err)
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrParseResponseBody, err)
		}

		return nil, fmt.Errorf("%w: server returned %v", ErrPerformRequest, body)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseResponseBody, err)
	}

	return body, nil
}
