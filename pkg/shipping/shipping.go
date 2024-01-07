package shipping

type (
	ShippingDataRequest struct {
		Shipper        ShipperInfo       `json:"shipper"         validate:"required"`
		Recipient      Address           `json:"recipient"       validate:"required"`
		Dispatchers    []DispatchersInfo `json:"dispatchers"     validate:"required"`
		SimulationType []int             `json:"simulation_type" validate:"required"`
	}

	DispatchersInfo struct {
		RegisteredNumber string   `json:"registered_number" validate:"required"`
		Zipcode          int64    `json:"zipcode"           validate:"required"`
		Volumes          []Volume `json:"volumes"`
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
		Type    int    `json:"type"`
		Country string `json:"country" validate:"required"`
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
		Carrier      CarrierInfo  `json:"carrier"`
		Service      string       `json:"service,omitempty"`
		FinalPrice   float64      `json:"final_price"`
		DeliveryTime DeliveryTime `json:"delivery_time"`
	}

	DeliveryTime struct {
		Days int64 `json:"days"`
	}
)
