package handler

import "frete-rapido-api/pkg/shipping"

type handler struct {
	shippingSvc shipping.Services
}

func New() (shipping.Handler, error) {
	h := &handler{}

	return h, nil
}
