package handler

import "shipping-calculator-api/pkg/shipping"

type Option func(*handler) error

func WithService(shippingSvc shipping.Services) Option {
	return func(h *handler) error {
		h.shippingSvc = shippingSvc
		return nil
	}
}
