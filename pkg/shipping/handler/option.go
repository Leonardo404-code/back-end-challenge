package handler

type Option func(*handler) error

func WithService() Option {
	return func(h *handler) error {
		return nil
	}
}
