package handler

import "frete-rapido-api/pkg/shipping"

type handler struct {
	shippingSvc shipping.Services
}

func New(opts ...Option) (shipping.Handler, error) {
	h := &handler{}

	for _, opt := range opts {
		if err := opt(h); err != nil {
			return nil, err
		}
	}

	return h, nil
}

func Must(shippingSvc shipping.Services) shipping.Handler {
	shipping, err := New(WithService(shippingSvc))
	if err != nil {
		panic(err)
	}

	return shipping
}
