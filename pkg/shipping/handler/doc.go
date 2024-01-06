package handler

type (
	QuotesResponseDoc struct {
		Carriers []Carrier
	}

	Carrier struct {
		Name     string `json:"name,omitempty"     example:"Expresso FR"`
		Service  string `json:"service,omitempty"  example:"Rodoviário"`
		Deadline string `json:"deadline,omitempty" example:"3"`
		Price    int    `json:"price,omitempty"    example:"17"`
	}

	ShippingDataDoc struct {
		Recipient Address
		Volumes   []Volume
	}

	Volume struct {
		Category      string  `json:"category,omitempty"       example:"7"`
		Amount        int     `json:"amount,omitempty"         example:"1"`
		UnitaryWeight int     `json:"unitary_weight,omitempty" example:"5"`
		Price         int     `json:"price,omitempty"          example:"350"`
		Height        float64 `json:"height,omitempty"         example:"0.2"`
		Width         float64 `json:"width,omitempty"          example:"0.2"`
		Length        float64 `json:"length,omitempty"         example:"0.2"`
	}

	Address struct {
		Zipcode string `json:"zipcode,omitempty" example:"01311000"`
	}
)
