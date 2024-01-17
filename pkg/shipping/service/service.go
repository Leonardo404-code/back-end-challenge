package service

import (
	"sync"

	"shipping-calculator-api/pkg/shipping"
)

type service struct {
	shippingRepo shipping.Repository
}

var (
	svc  shipping.Services
	once sync.Once
)

func Must(shippingRepo shipping.Repository) shipping.Services {
	once.Do(func() {
		svc = &service{
			shippingRepo: shippingRepo,
		}
	})

	return svc
}
