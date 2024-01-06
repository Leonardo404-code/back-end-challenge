package shipping

type (
	ShippingDataRequest struct {
		Recipient Address
		Volumes   []Volume
	}

	Volume struct {
		Category      string  `json:"category,omitempty"`
		Amount        int     `json:"amount,omitempty"`
		UnitaryWeight int     `json:"unitary_weight,omitempty"`
		Price         int     `json:"price,omitempty"`
		Height        float64 `json:"height,omitempty"`
		Width         float64 `json:"width,omitempty"`
		Length        float64 `json:"length,omitempty"`
	}

	Address struct {
		Zipcode string `json:"zipcode,omitempty"`
	}

	ShippingDataResponse struct{}
)
