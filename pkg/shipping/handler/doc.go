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
		Recipient Address
		Volumes   []Volume
	}

	Volume struct {
		Category      string  `json:"category"       example:"7"`
		Amount        int     `json:"amount"         example:"1"`
		UnitaryWeight int     `json:"unitary_weight" example:"5"`
		Price         int     `json:"price"          example:"350"`
		Height        float64 `json:"height"         example:"0.2"`
		Width         float64 `json:"width"          example:"0.2"`
		Length        float64 `json:"length"         example:"0.2"`
	}

	Address struct {
		Zipcode string `json:"zipcode,omitempty" example:"01311000"`
	}

	MetricsResposneDoc struct {
		Carriers    []Carrier `json:"carriers"`
		LowestPrice float64   `json:"lowest_price" example:"15.5"`
		HigherPrice float64   `json:"higher_price" example:"50.80"`
	}
)
