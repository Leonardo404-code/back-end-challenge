package shipping

type Repository interface {
	Create(carrier *CarrierInfo) error
}
