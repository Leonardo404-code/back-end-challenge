package repository

type ShippingDBModel struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int64   `json:"deadline"`
	Price    float64 `json:"price"`
}
