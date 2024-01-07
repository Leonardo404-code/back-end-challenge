package shipping

type Services interface {
	Quotes(shippingData *ShippingDataRequest) (*ShippingDataResponse, error)
}
