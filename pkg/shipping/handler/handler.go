package handler

import "frete-rapido-api/pkg/shipping"

type handler struct{}

func New() (shipping.Handler, error) {
	h := &handler{}

	return h, nil
}

func Must() shipping.Handler {
	shipping, err := New()
	if err != nil {
		panic(err)
	}

	return shipping
}
