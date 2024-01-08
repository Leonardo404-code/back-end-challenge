package shipping

type Services interface {
	Quotes(shippingData *ShippingDataRequest) (*ShippingDataResponse, error)
	Metrics(filter *Filter) (*MetricsResponse, error)
}
