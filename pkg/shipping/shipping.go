package shipping

type (
	ShippingDataRequest struct {
		Shipper        ShipperInfo       `json:"shipper"`
		Dispatchers    []DispatchersInfo `json:"dispatchers"`
		SimulationType []int             `json:"simulation_type"`
		Recipient      Address           `json:"recipient"`
	}

	DispatchersInfo struct {
		RegisteredNumber string   `json:"registered_number" validate:"required"`
		Volumes          []Volume `json:"volumes"`
		Zipcode          int64    `json:"zipcode"           validate:"required"`
	}

	Volume struct {
		Category      string  `json:"category"       validate:"required"`
		Amount        int64   `json:"amount"         validate:"required"`
		UnitaryWeight float64 `json:"unitary_weight" validate:"required"`
		UnitaryPrice  float64 `json:"unitary_price"  validate:"required"`
		Price         float64 `json:"price"          validate:"required"`
		Height        float64 `json:"height"         validate:"required"`
		Width         float64 `json:"width"          validate:"required"`
		Length        float64 `json:"length"         validate:"required"`
	}

	Address struct {
		Country string `json:"country" validate:"required"`
		Type    int    `json:"type"`
		Zipcode int64  `json:"zipcode" validate:"required"`
	}

	ShipperInfo struct {
		RegisteredNumber string `json:"registered_number" validate:"required"`
		Token            string `json:"token"             validate:"required"`
		PlatformCode     string `json:"platform_code"     validate:"required"`
	}

	ShippingDataResponse struct {
		Carrier []CarrierInfo `json:"carrier"`
	}

	CarrierInfo struct {
		Name     string  `json:"name"`
		Service  string  `json:"service"`
		Deadline int64   `json:"deadline"`
		Price    float64 `json:"price"`
	}

	QuotesData struct {
		Dispatcher []Dispatcher `json:"dispatchers"`
	}

	Dispatcher struct {
		Offers []Offers `json:"offers"`
	}

	Offers struct {
		Service      string       `json:"service,omitempty"`
		Carrier      CarrierInfo  `json:"carrier"`
		FinalPrice   float64      `json:"final_price"`
		DeliveryTime DeliveryTime `json:"delivery_time"`
	}

	DeliveryTime struct {
		Days int64 `json:"days"`
	}

	CarrierDBModel struct {
		ID       string  `json:"id"`
		Name     string  `json:"name"`
		Service  string  `json:"service"`
		Deadline int64   `json:"deadline"`
		Price    float64 `json:"price"`
	}

	MetricsResponse struct {
		Carriers    map[string]CarrierMetrics `json:"carriers"`
		LowestPrice float64                   `json:"lowest_price"`
		HigherPrice float64                   `json:"higher_price"`
	}

	CarrierMetrics struct {
		Total        float64 `json:"total"`
		AveragePrice float64 `json:"average_price"`
		Results      int64   `json:"results"`
	}
)
