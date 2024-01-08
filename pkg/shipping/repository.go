package shipping

type Repository interface {
	Create(carrier *CarrierInfo) error
	Search(filter *Filter) ([]*CarrierDBModel, error)
}
