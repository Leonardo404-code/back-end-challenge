package handler

type (
	QuotesResponseDoc struct {
		Carriers []Carrier
	}

	Carrier struct {
		Name     string `json:"name"     example:"Expresso FR"`
		Service  string `json:"service"  example:"Normal"`
		Deadline string `json:"deadline" example:"3"`
		Price    int    `json:"price"    example:"17"`
	}

	ShippingDataDoc struct {
		Shipper        ShipperDoc
		Recipient      AddressDoc
		Dispatchers    []DispatchersDoc
		SimulationType []int `json:"simulation_type"`
	}

	ShipperDoc struct {
		RegisteredNumber string `json:"registered_number" example:"12345"`
		Token            string `json:"token"             example:"19ebbe852f2f48be9713e802522783ee"`
		PlatformCode     string `json:"category"          example:"0ScdOLqC3r"`
	}

	DispatchersDoc struct {
		RegisteredNumber string `json:"registered_number" example:"81321255683644951537"`
		Zipcode          int    `json:"zipcode"           example:"1324553"`
		Volumes          []Volume
	}

	Volume struct {
		Category      string  `json:"category"       example:"7"`
		Amount        int     `json:"amount"         example:"1"`
		UnitaryWeight float64 `json:"unitary_weight" example:"5.2"`
		UnitaryPrice  float64 `json:"unitary_price"  example:"5.2"`
		Price         int     `json:"price"          example:"350"`
		Height        float64 `json:"height"         example:"0.2"`
		Width         float64 `json:"width"          example:"0.2"`
		Length        float64 `json:"length"         example:"0.2"`
	}

	AddressDoc struct {
		Type    int    `json:"type"    example:"0"`
		Country string `json:"country" example:"BRA"`
		Zipcode int    `json:"zipcode" example:"1324553"`
	}

	MetricsResposneDoc struct {
		Carriers    map[string]CarrierMetricsDoc `json:"carriers"`
		LowestPrice float64                      `json:"lowest_price" example:"15.5"`
		HigherPrice float64                      `json:"higher_price" example:"55.50"`
	}

	CarrierMetricsDoc struct {
		Total        float64 `json:"total"         example:"55.50"`
		AveragePrice float64 `json:"average_price" example:"18.50"`
		Results      int64   `json:"results"       example:"3"`
	}
)
