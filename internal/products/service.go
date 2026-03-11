package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

// ~ so whatever defines with in the service interface it have to satisfy that out
func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}
